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

type RabbitConsumer interface {
	Consume(handler MsgHandler)
}

type Consumer struct {
	// 原生配置
	conf ConsumerConf

	// 连接
	conn *amqp.Connection
	// 信道
	ch *amqp.Channel
	// 关闭通知信道
	closeNotifyChan chan *amqp.Error
	// 消息通道
	deliveries <-chan amqp.Delivery

	// 停止信道
	stopChan chan struct{}
	// 结束信道
	done chan struct{}
}

const DefaultBatchPullCount = 30

func NewConsumer(conf ConsumerConf) *Consumer {
	if conf.BatchPullCount == 0 {
		conf.BatchPullCount = DefaultBatchPullCount
	}

	return &Consumer{
		conf:     conf,
		stopChan: make(chan struct{}),
		done:     make(chan struct{}),
	}
}

type MsgHandler func(delivery amqp.Delivery)

// Watch 监听连接断开, 然后重连.
func (c *Consumer) Watch() {
	oldConn := c.conn
	oldCh := c.ch

watchConsumerLoop:
	for {
		select {
		case <-c.closeNotifyChan:
			if err := c.Setup(); err != nil {
				log.Error("rabbitmq consumer reconnect failed", zap.Error(err))
				time.Sleep(enums.ThreeSecD)
			} else {
				oldCh.Cancel(c.conf.ConsumerTag, true)
				oldConn.Close()
				oldConn, oldCh = c.conn, c.ch
				log.Info("rabbitmq consumer reconnect success")
			}
		case <-c.stopChan:
			break watchConsumerLoop
		default:
			time.Sleep(enums.OneSecD)
		}
	}

	c.done <- struct{}{}
}

func (c *Consumer) Setup() error {
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

	err = c.fetchDeliveries(ch)
	if err != nil {
		return err
	}

	c.conn, c.ch = conn, ch
	c.closeNotifyChan = conn.NotifyClose(make(chan *amqp.Error))

	return nil
}

func (c *Consumer) Start() error {
	err := c.Setup()
	if err != nil {
		return err
	}

	go c.Watch()
	return nil
}

func (c *Consumer) fetchDeliveries(ch *amqp.Channel) error {
	// 设置每次拉取消息的数量, 默认是30条
	if err := ch.Qos(c.conf.BatchPullCount, 0, false); err != nil {
		return err
	}

	deliveries, err := ch.Consume(
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

	c.deliveries = deliveries
	return nil
}

// Consume 消费消息.
func (c *Consumer) Consume(handler MsgHandler) {
	for {
		for delivery := range c.deliveries {
			handler(delivery)
		}
		time.Sleep(enums.FiveSecD)
	}
}

// Close 关闭.
func (c *Consumer) Close() error {
	c.stopChan <- struct{}{}
	<-c.done

	if c.ch != nil {
		if err := c.ch.Cancel(c.conf.ConsumerTag, true); err != nil {
			log.Error("close rabbitmq channel consumer error", zap.Error(err))
			return CancelChannelErr
		}
	}

	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			log.Error("close rabbitmq connection consumer error", zap.Error(err))
			return CloseConnErr
		}
	}

	log.Info("close rabbitmq consumer success")

	return nil
}
