package rabbitmq

import (
	"time"

	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type ConsumerConf struct {
	Url            string
	RoutingKeys    []string
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

const DefaultBatchPullCount = 30

func NewConsumer(conf ConsumerConf) *Consumer {
	if conf.BatchPullCount == 0 {
		conf.BatchPullCount = DefaultBatchPullCount
	}

	return &Consumer{
		conf: conf,
		done: make(chan struct{}),
	}
}

type MsgHandler func(delivery amqp.Delivery)

// ReConnect 重连.
func (c *Consumer) ReConnect() {
	closeChan := c.conn.NotifyClose(make(chan *amqp.Error))
	oldConn := c.conn
	oldCh := c.ch
	for {
		select {
		case <-closeChan:
			if err := c.Start(); err != nil {
				log.Error("reconnect rabbitmq failed", zap.Error(err))
				time.Sleep(enums.ThreeSecD)
			} else {
				oldCh.Close()
				oldConn.Close()
			}
		case <-c.done:
			return
		default:
			time.Sleep(enums.OneSecD)
		}
	}
}

// Start 开启消费者.
func (c *Consumer) Start() error {
	c.conf.Exchange.checkAndSet()

	conn, err := Connect(c.conf.Url)
	if err != nil {
		return err
	}

	// 获取信道
	ch, channelErr := conn.Channel()
	if channelErr != nil {
		return channelErr
	}

	// 声明交换机
	if err = ExchangeDeclare(ch, c.conf.Exchange); err != nil {
		return err
	}

	// 声明队列
	err = QueueDeclare(ch, c.conf.Queue)
	if err != nil {
		return err
	}

	// 交换机与队列绑定
	broadcast := c.conf.Exchange.ExType == Fanout
	if err = BindExchangeQueue(
		ch,
		c.conf.Queue.Name,
		c.conf.Exchange.Name,
		c.conf.RoutingKeys,
		broadcast); err != nil {
		return err
	}

	// 监听连接断开, 然后重连
	c.conn, c.ch = conn, ch
	go func() {
		c.ReConnect()
	}()

	return nil
}

// Consume 消费消息.
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

	for {
		select {
		case d, open := <-deliveries:
			if !open {
				return nil
			}
			handler(d)
		case <-c.done:
			return nil
		default:
			time.Sleep(time.Second)
		}
	}
}

// Close 关闭.
func (c *Consumer) Close() error {
	c.done <- struct{}{}
	log.Info("close rabbitmq consumer")
	if err := c.ch.Close(); err != nil {
		log.Error("close rabbitmq channel consumer error", zap.Error(err))
		return CancelChannelErr
	}

	if err := c.conn.Close(); err != nil {
		log.Error("close rabbitmq connection consumer error", zap.Error(err))
		return CloseConnErr
	}

	return nil
}
