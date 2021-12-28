package client

import (
	"github.com/gin-gonic/gin"
)

type Setting struct {
	Host            string            `json:"host" yaml:"host"`                       //地址
	Port            int               `json:"port" yaml:"port"`                       //端口, 默认是9000
	Language        string            `json:"language" yaml:"language"`               //校验错误返回的语言
	ShutDownTimeout int               `json:"shutDownTimeout" yaml:"shutDownTimeout"` //优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒， 默认5s
	Application     string            `json:"application" yaml:"application"`         //应用名
	Debug           bool              `json:"debug" yaml:"debug"`                     //debug, 默认false
	Swagger         bool              `json:"swagger" yaml:"swagger"`                 //是否启动swagger, 默认false
	LogConfig       Log               `json:"log" yaml:"log"`                         //日志配置
	UseRtePool      bool              `json:"useRtePool" yaml:"useRtePool"`           //使用协程池, 默认为true
	Middlewares     []gin.HandlerFunc //中间件
	OnStartUp       []func() error    //项目启动前执行函数
	OnShutDown      []func() error    //项目关闭前执行函数
}

//Arrange 处理零值及无效字段为默认值
func (s *Setting) Arrange() {
	if s.Host == "" {
		s.Host = "127.0.0.1"
	}

	if s.Port <= 0 {
		s.Port = 9000
	}

	if s.ShutDownTimeout <= 0 {
		s.ShutDownTimeout = 5
	}
}
