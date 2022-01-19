package setting

import (
	"github.com/panjf2000/ants/v2"
)

// Routine 协程池配置
type Routine struct {
	// MaxWorkerCount 最大worker数量, 默认是5000
	MaxWorkerCount int `json:"maxWorkerCount" yaml:"maxWorkerCount"`

	// WaitTimeOut 阻塞后, 等待超时时间, 小于或等于0则没有超时时间, 单位: 秒
	//WaitTimeOut int `json:"maxRetries" yaml:"maxRetries"`

	// TaskTimeOut 执行任务超时时间, 小于或等于0则没有超时时间, 单位: 秒
	//TaskTimeOut int `json:"taskTimeOut" yaml:"taskTimeOut"`

	// FreeMaxLifetime 协程最大闲置时间, 默认是1小时, 单位: 秒
	FreeMaxLifetime int `json:"freeMaxLifetime" yaml:"freeMaxLifetime"`

	// PanicHandler panic处理器
	PanicHandler func(interface{})

	// Logger 日志处理器
	Logger ants.Logger
}
