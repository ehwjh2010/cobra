package client

import "time"

type Server struct {
	Host            string        `yaml:"host" json:"host"`                       //服务器IP
	Port            int           `yaml:"port" json:"port"`                       //服务器端口
	ShutDownTimeout time.Duration `yaml:"shutDownTimeout" json:"shutDownTimeout"` //优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒
}
