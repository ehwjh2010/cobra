package rabbitmq

import (
	"context"
	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/helper/basic/collection"
	"github.com/ehwjh2010/viper/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

type ProducerConf struct {
	Url      string
	Exchange Exchange
	Queue    Queue
}

type Producer struct {
	conf ProducerConf
	conn *amqp.Connection
	ch   *amqp.Channel
	done chan struct{}
}

func NewProducer(conf ProducerConf) *Producer {
	return &Producer{
		conf: conf,
		done: make(chan struct{}),
	}
}

// ReConnect 重连.
func (p *Producer) ReConnect() {
	for {
		select {
		case <-p.conn.NotifyClose(make(chan *amqp.Error)):
			oldConn := p.conn
			oldCh := p.ch
			if err := p.Start(); err != nil {
				oldCh.Close()
				oldConn.Close()
			}
		case <-p.done:
			return
		default:
			time.Sleep(enums.ThreeSecD)
		}
	}
}

func (p *Producer) Start() error {
	p.conf.Exchange.checkAndSet()

	conn, err := Connect(p.conf.Url)
	if err != nil {
		return err
	}

	// 获取信道
	ch, channelErr := conn.Channel()
	if channelErr != nil {
		return channelErr
	}

	// 声明交换机
	if err = ExchangeDeclare(ch, p.conf.Exchange); err != nil {
		return err
	}

	// 声明队列
	if err = QueueDeclare(ch, p.conf.Queue); err != nil {
		return err
	}

	// 监听连接断开, 然后重连
	p.ch, p.conn = ch, conn
	go func() {
		p.ReConnect()
	}()

	return nil
}

// SendMsg 发送消息.
func (p *Producer) SendMsg(ctx context.Context, key string, body []byte) error {
	if collection.IsEmptyBytes(body) {
		return nil
	}

	err := p.ch.PublishWithContext(
		ctx,
		p.conf.Exchange.Name,
		key,
		false, // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
		false, // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain", // 消息内容的类型,
			Body:         body,
		})

	return err
}

// Close 关闭.
func (p *Producer) Close() error {
	p.done <- struct{}{}

	if err := p.ch.Close(); err != nil {
		log.Error("rabbitmq producer channel close failed", zap.Error(err))
		return CancelChannelErr
	}

	if err := p.conn.Close(); err != nil {
		log.Error("rabbitmq producer connection close failed", zap.Error(err))
		return CloseConnErr
	}

	return nil
}
