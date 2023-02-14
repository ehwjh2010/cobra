package server

import (
	"github.com/ehwjh2010/viper/component/routine"
	"github.com/ehwjh2010/viper/log"
	"github.com/gin-gonic/gin"
)

type Setting struct {
	Host            string            `json:"host" yaml:"host"`                       // 地址
	Port            int               `json:"port" yaml:"port"`                       // 端口, 默认是9000
	Language        string            `json:"language" yaml:"language"`               // 校验错误返回的语言
	ShutDownTimeout int               `json:"shutDownTimeout" yaml:"shutDownTimeout"` // 优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒， 默认3s
	Application     string            `json:"application" yaml:"application"`         // 应用名
	Debug           bool              `json:"debug" yaml:"debug"`                     // debug, 默认false
	Swagger         bool              `json:"swagger" yaml:"swagger"`                 // 是否启动swagger, 默认false
	LogConfig       log.ZapLogReq     `json:"log" yaml:"log"`                         // 日志配置
	EnableRtPool    bool              `json:"enableRtPool" yaml:"enableRtPool"`       // 启用协程池, 默认是false
	Routine         routine.Routine   `json:"routine" yaml:"routine"`                 // 协程池配置
	Middlewares     []gin.HandlerFunc // 中间件
	OnHookFunc                        // server勾子函数
	GrpcConf        *GraceGrpc        // grpc配置
}

// Arrange 处理零值及无效字段为默认值.
func (s *Setting) Arrange() {
	if s.Host == "" {
		s.Host = "127.0.0.1"
	}

	if s.Port <= 0 {
		s.Port = 12345
	}

	if s.ShutDownTimeout <= 0 {
		s.ShutDownTimeout = 3
	}
}
