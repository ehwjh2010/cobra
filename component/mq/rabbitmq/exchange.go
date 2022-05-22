package rabbitmq

import "github.com/ehwjh2010/viper/helper/basic/str"

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

// CheckAndSet 处理默认值
func (e *Exchange) CheckAndSet() error {
	if str.IsEmpty(e.ExType) {
		e.ExType = Direct
	}

	return nil
}
