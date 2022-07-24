package rabbitmq

import (
	"github.com/ehwjh2010/viper/helper/basic/collection"
	"time"

	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type ProducerConf struct {
	Url        string
	RoutingKey string
	Exchange   Exchange
	Queue      Queue
}

type Producer struct {
	conf ProducerConf
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewProducer(conf ProducerConf) *Producer {
	return &Producer{
		conf: conf,
	}
}

// ReConnect 重连
func (p *Producer) ReConnect() {

	for {
		if err := p.Start(); err != nil {
			log.Error("rabbitmq reconnect failed", zap.Error(err))
			time.Sleep(enums.ThreeSecD)
		} else {
			break
		}
	}

}

func (p *Producer) Start() error {
	p.conf.Exchange.checkAndSet()

	conn, err := Connect(p.conf.Url)
	if err != nil {
		return err
	}

	// 监听连接断开, 然后重连
	go func() {
		for {
			<-p.conn.NotifyClose(make(chan *amqp.Error))
			p.ReConnect()
		}
	}()

	p.conn = conn

	// 获取信道
	if p.ch, err = p.conn.Channel(); err != nil {
		return err
	}

	// 声明交换机
	if err = ExchangeDeclare(p.ch, p.conf.Exchange); err != nil {
		return err
	}

	// 声明队列
	if err = QueueDeclare(p.ch, p.conf.Queue); err != nil {
		return err
	}

	// 交换机与队列绑定
	broadcast := p.conf.Exchange.ExType == Fanout
	if err = BindExchangeQueue(p.ch, p.conf.Queue.Name, p.conf.Exchange.Name, p.conf.RoutingKey, broadcast); err != nil {
		return err
	}

	return nil
}

// SendMsg 发送消息
func (p *Producer) SendMsg(body []byte) error {
	if collection.IsEmptyBytes(body) {
		return nil
	}

	err := p.ch.Publish(
		p.conf.Exchange.Name,
		p.conf.RoutingKey,
		false, // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
		false, // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "text/plain", // 消息内容的类型,
			Body:         body,
		})

	return err
}

// Close 关闭
func (c *Producer) Close() error {
	if err := c.ch.Close(); err != nil {
		return CancelChannelErr
	}

	if err := c.conn.Close(); err != nil {
		return CloseConnErr
	}

	return nil
}
