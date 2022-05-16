package rabbitmq

import (
	"errors"
	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/ehwjh2010/viper/log"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var (
	EmptyConnection = errors.New("empty connection get channel")
	EmptyExchange   = errors.New("empty exchange")
	EmptyQueue      = errors.New("empty queue")
)

type HandlerMsg int

const (
	Direct  = "direct"
	Fanout  = "fanout"
	Topic   = "topic"
	Headers = "headers"
)

const (
	Send HandlerMsg = iota + 1
	Recv
)

type RabbitMQ struct {
	Url        string    `json:"url,omitempty" yaml:"url"`
	RoutingKey string    `json:"routingKey,omitempty" yaml:"routingKey"`
	Exchange   *Exchange `json:"exchange,omitempty" yaml:"exchange"`
	Queue      *Queue    `json:"queue,omitempty" yaml:"queue"`

	connection  *amqp.Connection
	sendChannel *amqp.Channel
	recvChannel *amqp.Channel
}

// Exchange 交换机
type Exchange struct {
	Name        string `json:"name" yaml:"name"`               // 名字
	ExType      string `json:"exType" yaml:"exType"`           // 类型
	Persistence bool   `json:"persistence" yaml:"persistence"` // 持久化
	AutoDeleted bool   `json:"autoDeleted" yaml:"autoDeleted"` // 自动删除
	Internal    bool   `json:"internal" yaml:"internal"`
	NoWait      bool   `json:"noWait" yaml:"noWait"`
	Arguments   map[string]interface{}
}

// CheckAndSet 处理默认值
func (e *Exchange) CheckAndSet() error {
	if str.IsEmpty(e.ExType) {
		e.ExType = Direct
	}

	if str.IsEmpty(e.Name) {
		return EmptyExchange
	}

	return nil
}

type Queue struct {
	Name        string // 名字
	Persistence bool   `json:"persistence" yaml:"persistence"` // 持久化
	AutoDeleted bool   `json:"autoDeleted" yaml:"autoDeleted"` // 自动删除
	Exclusive   bool   `json:"exclusive" yaml:"exclusive"`
	NoWait      bool   `json:"noWait" yaml:"noWait"`
	Arguments   map[string]interface{}
}

// CheckAndSet 处理默认值
func (q *Queue) CheckAndSet() error {
	if str.IsEmpty(q.Name) {
		return EmptyQueue
	}

	return nil
}

func (r *RabbitMQ) getChannel(t HandlerMsg) *amqp.Channel {
	var ch *amqp.Channel
	switch t {
	case Send:
		ch = r.sendChannel
	case Recv:
		ch = r.recvChannel
	default:
		panic("not support")
	}

	return ch
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
		log.Error("open rabbitmq channel failed", zap.Error(err), zap.String("Url", r.Url))
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

	sendCh, err := r.connection.Channel()
	if err != nil {
		log.Error("open rabbitmq send channel failed", zap.Error(err), zap.String("Url", r.Url))
		return err
	}
	r.sendChannel = sendCh

	recvCh, err := r.connection.Channel()
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
	return nil
}

// Start 开始
func (r *RabbitMQ) Start() error {
	if r.Exchange == nil {
		return EmptyExchange
	}

	if err := r.Exchange.CheckAndSet(); err != nil {
		return err
	}

	if r.Queue == nil {
		return EmptyQueue
	}

	if err := r.Queue.CheckAndSet(); err != nil {
		return err
	}

	if err := r.Connect(); err != nil {
		return err
	}

	if err := r.OpenChannel(); err != nil {
		return err
	}

	return nil
}

// ExchangeDeclare 声明交换机
func (r *RabbitMQ) ExchangeDeclare(ch *amqp.Channel) error {
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
	if !str.IsEmpty(r.RoutingKey) && !str.IsEmpty(r.Exchange.Name) {
		if err := ch.QueueBind(r.Queue.Name, r.RoutingKey, r.Exchange.Name, false, nil); err != nil {
			return err
		}
	}

	return nil
}

// SendMsg 发送消息
func (r *RabbitMQ) SendMsg(content []byte) error {
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

	// TODO ZZZZZZZZZZZZ
	return nil
}
