package middleware

import (
	"github.com/ehwjh2010/cobra/enum"
	"github.com/gin-gonic/gin"
)

const (
	UTC   = false
	Stack = true
)

//MiddleConfig 中间件配置
type MiddleConfig struct {
	//TimeFormat 时间格式
	TimeFormat string `yaml:"timeFormat" json:"timeFormat"`

	//UTC 是否使用UTC时间, 否则使用本地时间
	UTC bool `yaml:"utc" json:"utc"`

	//SkipPaths  跳过url
	SkipPaths []string `yaml:"skipPaths" json:"skipPaths"`

	//Stack 是否打印堆栈信息
	Stack bool `yaml:"stack" json:"stack"`
}

func NewMiddleConfig(args ...MiddleConfigOption) (config *MiddleConfig) {
	config = &MiddleConfig{
		TimeFormat: enum.DefaultTimePattern,
		UTC:        UTC,
		SkipPaths:  nil,
		Stack:      Stack,
	}

	for _, arg := range args {
		arg(config)
	}

	return config
}

type MiddleConfigOption func(middleConfig *MiddleConfig)

func MiddleConfigWithTimeFormat(timeFormat string) MiddleConfigOption {
	return func(middleConfig *MiddleConfig) {
		middleConfig.TimeFormat = timeFormat
	}
}

func MiddleConfigWithUTC(utc bool) MiddleConfigOption {
	return func(middleConfig *MiddleConfig) {
		middleConfig.UTC = utc
	}
}

func MiddleConfigWithSkipPath(skipPath []string) MiddleConfigOption {
	return func(middleConfig *MiddleConfig) {
		middleConfig.SkipPaths = skipPath
	}
}

func MiddleConfigWithStack(stack bool) MiddleConfigOption {
	return func(middleConfig *MiddleConfig) {
		middleConfig.Stack = stack
	}
}

type MidFunc func(config *MiddleConfig) gin.HandlerFunc

var middlewares = []MidFunc{
	GinZap,
	RecoveryWithZap,
}

func UseMiddles(handler *gin.Engine, config *MiddleConfig) {
	if len(middlewares) == 0 {
		return
	}

	for _, middleware := range middlewares {
		handler.Use(middleware(config))
	}
}
