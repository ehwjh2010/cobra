package rabbitmq

import (
	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

type ConsumerConf struct {
	Url            string
	RoutingKey     string
	Exchange       Exchange
	Queue          Queue
	BatchPullCount int
	ConsumerTag    string
	AutoAck        bool
	Exclusive      bool
	NoLocal        bool
	NoWait         bool
	Args           amqp.Table
}

type Consumer struct {
	conf ConsumerConf

	conn *amqp.Connection
	ch   *amqp.Channel
	done chan struct{}
}

func NewConsumer(conf ConsumerConf) *Consumer {
	return &Consumer{
		conf: conf,
		done: make(chan struct{}),
	}
}

type MsgHandler func(delivery amqp.Delivery)

// ReConnect 重连
func (c *Consumer) ReConnect() {

	for {
		if err := c.Start(); err != nil {
			log.Error("rabbitmq reconnect failed", zap.Error(err))
			time.Sleep(enums.ThreeSecD)
		} else {
			break
		}
	}

}

// Start 开启消费者
func (c *Consumer) Start() error {
	c.conf.Exchange.checkAndSet()

	conn, err := Connect(c.conf.Url)
	if err != nil {
		return err
	}

	c.conn = conn

	// 监听连接断开, 然后重连
	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error))
		c.ReConnect()
	}()

	// 获取信道
	if c.ch, err = c.conn.Channel(); err != nil {
		return err
	}

	// 声明交换机
	if err = ExchangeDeclare(c.ch, c.conf.Exchange); err != nil {
		return err
	}

	// 声明队列
	if err = QueueDeclare(c.ch, c.conf.Queue); err != nil {
		return err
	}

	// 交换机与队列绑定
	broadcast := c.conf.Exchange.ExType == Fanout
	if err = BindExchangeQueue(c.ch, c.conf.Queue.Name, c.conf.Exchange.Name, c.conf.RoutingKey, broadcast); err != nil {
		return err
	}

	return nil
}

// Consume 消费消息
func (c *Consumer) Consume(handler MsgHandler) error {
	// 设置每次拉取消息的数量, 默认是30条
	if err := c.ch.Qos(c.conf.BatchPullCount, 0, false); err != nil {
		return err
	}

	deliveries, err := c.ch.Consume(
		c.conf.Queue.Name,
		c.conf.ConsumerTag,
		c.conf.AutoAck,
		c.conf.Exclusive,
		c.conf.NoLocal,
		c.conf.NoWait,
		c.conf.Args)

	if err != nil {
		return err
	}

	log.Info("get deliveries success")

	for delivery := range deliveries {
		handler(delivery)
	}

	c.done <- struct{}{}
	return nil
}

// Close 关闭
func (c *Consumer) Close() error {
	log.Info("close rabbitmq consumer")
	if err := c.ch.Cancel(c.conf.ConsumerTag, false); err != nil {
		log.Error("cancel consumer", zap.Error(err))
		return CancelChannelErr
	}

	if err := c.conn.Close(); err != nil {
		log.Error("close conn", zap.Error(err))
		return CloseConnErr
	}

	// wait for handle() to exit
	<-c.done
	return nil
}
