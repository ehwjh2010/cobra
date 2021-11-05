package setting

import (
	"ginLearn/utils"
	"time"
)

type Config struct {
	Env             string             `yaml:"env" json:"env"`                         //环境标识
	Application     string             `yaml:"application" json:"application"`         //应用名
	ServerPort      uint               `yaml:"serverPort" json:"serverPort"`           //服务器端口
	ShutDownTimeout time.Duration      `yaml:"shutDownTimeout" json:"shutDownTimeout"` //#优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒
	LogConfig       *utils.LogConfig   `yaml:"log" json:"log"`
	DBConfig        *utils.DBConfig    `yaml:"db" json:"db"`
	RedisConfig     *utils.CacheConfig `yaml:"cache" json:"cache"`
}
