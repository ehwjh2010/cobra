package settings

type LogHandlerFunc func(string, ...interface{})

func (d LogHandlerFunc) Printf(format string, args ...interface{}) {
	d(format, args...)
}

// Routine 协程池配置
type Routine struct {
	// MaxWorkerCount 最大worker数量, 默认是10
	MaxWorkerCount int `json:"maxWorkerCount" yaml:"maxWorkerCount"`

	// FreeMaxLifetime 协程最大闲置时间, 默认是20分钟, 单位: 秒
	FreeMaxLifetime int `json:"freeMaxLifetime" yaml:"freeMaxLifetime"`

	// PanicHandler panic处理器
	PanicHandler func(interface{})

	// Logger 日志处理器
	Logger LogHandlerFunc
}
