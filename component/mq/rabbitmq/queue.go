package rabbitmq

import "github.com/ehwjh2010/viper/helper/basic/str"

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
