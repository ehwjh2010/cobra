package rabbitmq

import (
	"errors"
	"github.com/ehwjh2010/viper/helper/basic/str"
)

const (
	Direct  = "direct"
	Fanout  = "fanout"
	Topic   = "topic"
	Headers = "headers"
)

var (
	EmptyConnection  = errors.New("empty connection get channel")
	EmptyQueue       = errors.New("empty queue")
	EmptyRoutingKey  = errors.New("empty routing key")
	CancelChannelErr = errors.New("consumer cancel failed")
	CloseConnErr     = errors.New("amqp connection close error")
)

// Exchange 交换机
type Exchange struct {
	Name        string                 `json:"name" yaml:"name"`               // 交换器名
	ExType      string                 `json:"exType" yaml:"exType"`           // 类型
	Persistence bool                   `json:"persistence" yaml:"persistence"` // 持久化
	AutoDeleted bool                   `json:"autoDeleted" yaml:"autoDeleted"` // 是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
	Internal    bool                   `json:"internal" yaml:"internal"`       // 设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
	NoWait      bool                   `json:"noWait" yaml:"noWait"`           // 是否为非阻塞
	Arguments   map[string]interface{} // 额外属性
}

// checkAndSet 处理默认值
func (e *Exchange) checkAndSet() {
	if str.IsEmpty(e.ExType) {
		e.ExType = Direct
	}
}

type Queue struct {
	Name        string                 // 队列名
	Persistence bool                   `json:"persistence" yaml:"persistence"` // 持久化
	AutoDeleted bool                   `json:"autoDeleted" yaml:"autoDeleted"` // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
	Exclusive   bool                   `json:"exclusive" yaml:"exclusive"`     // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
	NoWait      bool                   `json:"noWait" yaml:"noWait"`           // 是否为非阻塞
	Arguments   map[string]interface{} // 额外属性
}

// checkAndSet 处理默认值
func (q *Queue) checkAndSet() error {
	if str.IsEmpty(q.Name) {
		return EmptyQueue
	}

	return nil
}
