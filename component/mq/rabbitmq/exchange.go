package rabbitmq

import "github.com/ehwjh2010/viper/helper/basic/str"

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

	return nil
}
