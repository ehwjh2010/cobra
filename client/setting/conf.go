package setting

import (
	"fmt"
	"time"
)

type Config struct {
	Env             string        `yaml:"env" json:"env"`                         //环境标识
	Application     string        `yaml:"application" json:"application"`         //应用名
	ServerPort      uint          `yaml:"serverPort" json:"serverPort"`           //服务器端口
	ShutDownTimeout time.Duration `yaml:"shutDownTimeout" json:"shutDownTimeout"` //#优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒
	LogConfig       *LogConfig    `yaml:"log" json:"log"`
	MysqlConfig     *MysqlConfig  `yaml:"mysql" json:"mysql"`
	RedisConfig     *RedisConfig  `yaml:"redis" json:"redis"`
}

type MysqlConfig struct {
	Host             string        `yaml:"host" json:"host"`                         //MySQL IP
	Port             uint          `yaml:"port" json:"port"`                         //MySQL 端口
	User             string        `yaml:"user" json:"user"`                         //用户名
	Password         string        `yaml:"password" json:"password"`                 //密码
	Database         string        `yaml:"database" json:"database"`                 //数据库名
	Location         string        `yaml:"location" json:"location"`                 //时区
	TablePrefix      string        `yaml:"tablePrefix" json:"tablePrefix"`           //表前缀
	SingularTable    bool          `yaml:"singularTable" json:"singularTable"`       //表复数禁用
	CreateBatchSize  int           `yaml:"createBatchSize" json:"createBatchSize"`   //批量创建数量
	EnableRawSQL     bool          `yaml:"enableRawSql" json:"enableRawSql"`         //打印原生SQL
	MaxFreeConnCount int           `yaml:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int           `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	FreeMaxLifetime  time.Duration `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   //闲置连接最大存活时间, 单位: 分钟
}

func (c *MysqlConfig) Dsn() string {
	uri := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s`,
		c.User, c.Password, c.Host, c.Port, c.Database, c.Location)
	return uri
}

type RedisConfig struct {
	Host             string        `yaml:"host" json:"host"`                         //Redis IP
	Port             uint          `yaml:"port" json:"port"`                         //Redis 端口
	Pwd              string        `yaml:"pwd" json:"pwd"`                           //密码
	MaxFreeConnCount int           `json:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int           `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	FreeMaxLifetime  time.Duration `json:"freeMaxLifetime" yaml:"freeMaxLifetime"`   //闲置连接存活的最大时间, 单位: 分钟
	Database         int           `yaml:"database" json:"database"`                 //数据库
	ConnectTimeout   time.Duration `yaml:"connectTimeout" json:"connectTimeout"`     //连接Redis超时时间, 单位: 秒
	ReadTimeout      time.Duration `yaml:"readTimeout" json:"readTimeout"`           //读取超时时间, 单位: 秒
	WriteTimeout     time.Duration `yaml:"writeTimeout" json:"writeTimeout"`         //写超时时间, 单位: 秒
}

type LogConfig struct {
	LogPath          string `yaml:"logPath" json:"logPath"`                   //日志文件目录
	Level            string `yaml:"level" json:"level"`                       //日志级别
	EnableLogConsole bool   `yaml:"enableLogConsole" json:"enableLogConsole"` //是否输出打终端
	AccessMethodRow  bool   `yaml:"accessMethodRow" json:"accessMethodRow"`   //是否打印方法以及所在行数
}
