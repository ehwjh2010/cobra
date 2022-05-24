package rabbitmq

import "github.com/ehwjh2010/viper/helper/basic/str"

type Queue struct {
	Name        string                 // 队列名
	Persistence bool                   `json:"persistence" yaml:"persistence"` // 持久化
	AutoDeleted bool                   `json:"autoDeleted" yaml:"autoDeleted"` // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
	Exclusive   bool                   `json:"exclusive" yaml:"exclusive"`     // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
	NoWait      bool                   `json:"noWait" yaml:"noWait"`           // 是否为非阻塞
	Arguments   map[string]interface{} // 额外属性
}

// CheckAndSet 处理默认值
func (q *Queue) CheckAndSet() error {
	if str.IsEmpty(q.Name) {
		return EmptyQueue
	}

	return nil
}
