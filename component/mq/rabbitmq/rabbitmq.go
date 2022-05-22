package rabbitmq

import (
	"github.com/ehwjh2010/viper/helper/wrapper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/verror"
)

type SendCallback func(body []byte) error

type RabbitMQ struct {
	Url                    string       `json:"url,omitempty" yaml:"url"`
	RoutingKey             string       `json:"routingKey,omitempty" yaml:"routingKey"`
	Exchange               *Exchange    `json:"exchange,omitempty" yaml:"exchange"`
	Queue                  *Queue       `json:"queue,omitempty" yaml:"queue"`
	SendSuccess            SendCallback // 发送成功后回调函数
	SendFail               SendCallback // 发送失败后回调函数
	SendTeardown           SendCallback // 不论失败还是成功, 发送后定触发的函数
	ConsumeSuccessCallback func([]byte, error) error

	// 消费配置
	PullBatchSize int                    `json:"pullBatchSize" yaml:"pullBatchSize"` // 批量拉取数量
	AutoAck       bool                   `json:"autoAck" yaml:"autoAck"`             // 自动提交
	Exclusive     bool                   `json:"exclusive" yaml:"exclusive"`         // 是否排他
	NoLocal       bool                   `json:"noLocal" yaml:"noLocal"`             // 是否接收只同一个连接中的消息，若为true，则只能接收别的conn中发送的消息
	NoWait        bool                   `json:"noWait" yaml:"noWait"`               // 非阻塞
	Args          map[string]interface{} // 额外参数

	// 内部参数
	startFlag   bool
	connection  *amqp.Connection
	sendChannel *amqp.Channel
	recvChannel *amqp.Channel
}

// Connect 连接
func (r *RabbitMQ) Connect() error {
	if r.connection != nil {
		return nil
	}

	connection, err := amqp.Dial(r.Url)
	if err != nil {
		log.Error("connect rabbitmq failed", zap.Error(err), zap.String("Url", r.Url))
		return err
	}
	r.connection = connection
	return nil
}

// Close 关闭mq链接
func (r *RabbitMQ) Close() error {
	if r.connection == nil {
		return nil
	}

	err := r.connection.Close()
	if err != nil {
		log.Error("close rabbitmq failed", zap.Error(err), zap.String("Url", r.Url))
		return err
	}
	return nil
}

// openChannel 开启Channel
func (r *RabbitMQ) openChannel() (*amqp.Channel, error) {

	if r.connection == nil {
		return nil, EmptyConnection
	}

	ch, err := r.connection.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// OpenChannel 开启Channel
func (r *RabbitMQ) OpenChannel() error {

	if err := r.Connect(); err != nil {
		return err
	}

	if r.sendChannel != nil && r.recvChannel != nil {
		return nil
	}

	sendCh, err := r.openChannel()
	if err != nil {
		log.Error("open rabbitmq send channel failed", zap.Error(err), zap.String("Url", r.Url))
		return err
	}
	r.sendChannel = sendCh

	recvCh, err := r.openChannel()
	if err != nil {
		log.Error("open rabbitmq recv channel failed", zap.Error(err), zap.String("Url", r.Url))
		return err
	}
	r.recvChannel = recvCh
	return nil
}

// CloseChannel 关闭Channel
func (r *RabbitMQ) closeChannel(ch *amqp.Channel) error {
	if ch == nil {
		return nil
	}

	return ch.Close()
}

// CloseChannel 关闭通道
func (r *RabbitMQ) CloseChannel() error {
	var multiErr verror.MultiErr

	multiErr.AddErr(r.closeChannel(r.recvChannel))
	multiErr.AddErr(r.closeChannel(r.sendChannel))

	err := multiErr.AsStdErr()

	return err
}

// Start 开始
func (r *RabbitMQ) Start() error {
	if r.Exchange == nil {
		return EmptyExchange
	}

	if r.Queue == nil {
		return EmptyQueue
	}

	if err := r.Exchange.CheckAndSet(); err != nil {
		return err
	}

	if err := r.Queue.CheckAndSet(); err != nil {
		return err
	}

	if r.SendSuccess == nil {
		r.SendSuccess = DefaultSuccessCallback
	}

	if r.SendFail == nil {
		r.SendFail = DefaultFailCallback
	}

	if r.PullBatchSize == 0 {
		r.PullBatchSize = DefaultPullBatchSize
	}

	if err := r.Connect(); err != nil {
		return err
	}

	if err := r.OpenChannel(); err != nil {
		return err
	}

	r.startFlag = true
	return nil
}

// ExchangeDeclare 声明交换机
func (r *RabbitMQ) ExchangeDeclare(ch *amqp.Channel) error {
	// 当交换机name为空, 则使用默认交换机, 不需要声明
	if !str.IsEmpty(r.Exchange.Name) {
		if err := ch.ExchangeDeclare(
			r.Exchange.Name,
			r.Exchange.ExType,
			r.Exchange.Persistence,
			r.Exchange.AutoDeleted,
			r.Exchange.Internal,
			r.Exchange.NoWait,
			r.Exchange.Arguments); err != nil {
			return err
		}
	}
	return nil
}

// QueueDeclare 声明队列
func (r *RabbitMQ) QueueDeclare(ch *amqp.Channel) error {
	if _, err := ch.QueueDeclare(
		r.Queue.Name,
		r.Queue.Persistence,
		r.Queue.AutoDeleted,
		r.Queue.Exclusive,
		r.Queue.NoWait,
		r.Queue.Arguments,
	); err != nil {
		return err
	}

	return nil
}

// BindExchangeQueue 交换机绑定队列
func (r *RabbitMQ) BindExchangeQueue(ch *amqp.Channel) error {
	if str.IsEmpty(r.RoutingKey) {
		return EmptyRoutingKeys
	}

	if err := ch.QueueBind(
		r.Queue.Name,    // 绑定的队列名称
		r.RoutingKey,    // bindkey 用于消息路由分发的key
		r.Exchange.Name, // 绑定的exchange名
		false,           // 是否阻塞
		nil,             // 额外属性
	); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) sendMsg(body []byte) error {
	// 声明交换机
	if err := r.ExchangeDeclare(r.sendChannel); err != nil {
		return err
	}

	// 声明队列
	if err := r.QueueDeclare(r.sendChannel); err != nil {
		return err
	}

	// 交换机绑定队列
	if err := r.BindExchangeQueue(r.sendChannel); err != nil {
		return err
	}

	err := r.sendChannel.Publish(
		r.Exchange.Name, // 交换器名
		r.RoutingKey,    // routing key
		false,           // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
		false,           // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
		amqp.Publishing{
			ContentType: "text/plain", // 消息内容的类型
			Body:        body,         // 消息内容
		})

	return err
}

// SendMsg 发送消息
func (r *RabbitMQ) SendMsg(body []byte) (bool, error) {
	if !r.startFlag {
		return false, ExecuteStartBeforeSendMsg
	}

	var multiErr verror.MultiErr

	err := r.sendMsg(body)
	ok := err == nil

	multiErr.AddErr(err)

	if err != nil {
		multiErr.AddErr(r.SendFail(body))
	} else {
		multiErr.AddErr(r.SendSuccess(body))
	}

	if r.SendTeardown != nil {
		multiErr.AddErr(r.SendTeardown(body))
	}

	err = multiErr.AsStdErr()

	return ok, err
}

// RecvMsg 接收消息
func (r *RabbitMQ) RecvMsg() (<-chan amqp.Delivery, error) {
	// 声明交换机
	if err := r.ExchangeDeclare(r.sendChannel); err != nil {
		return nil, err
	}

	// 声明队列
	if err := r.QueueDeclare(r.sendChannel); err != nil {
		return nil, err
	}

	// 交换机绑定队列
	if err := r.BindExchangeQueue(r.sendChannel); err != nil {
		return nil, err
	}

	if err := r.recvChannel.Qos(r.PullBatchSize, 0, false); err != nil {
		return nil, err
	}

	c, err := r.recvChannel.Consume(r.Queue.Name, "", r.AutoAck, r.Exclusive, r.NoLocal, r.NoWait, r.Args)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *RabbitMQ) consumeMsg(consumer Consumer) error {

	c, err := r.RecvMsg()
	if err != nil {
		return err
	}

	for d := range c {

		err := func() error {
			defer wrapper.PanicHandler()
			return consumer.Consume(d.Body)
		}()

		if err != nil {
			func() {
				defer wrapper.PanicHandler()
				consumer.ConsumeFailConsumeCallback(d.MessageId, d.Body, err)
			}()
			continue
		} else {
			if !r.AutoAck {
				err := d.Ack(false)
				if err != nil {
					log.Error("ack msg err", zap.ByteString("body", d.Body), zap.String("messageId", d.MessageId))
				}
			}
			consumer.ConsumeSuccessCallback(d.MessageId, d.Body)
		}
	}

	return nil
}

// ConsumeMsg 消费消息
func (r *RabbitMQ) ConsumeMsg(consumer Consumer) {
	for {
		if err := r.consumeMsg(consumer); err != nil {
			log.Error("consume msg failed", zap.Error(err))
		}
	}
}
