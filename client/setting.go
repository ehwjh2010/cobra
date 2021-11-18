package client

import (
	"fmt"
	"time"
)

type Config struct {
	Env             string        `yaml:"env" json:"env"`                         //环境标识
	Application     string        `yaml:"application" json:"application"`         //应用名
	ServerPort      uint          `yaml:"serverPort" json:"serverPort"`           //服务器端口
	ShutDownTimeout time.Duration `yaml:"shutDownTimeout" json:"shutDownTimeout"` //优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒
	LogConfig       *LogConfig    `yaml:"log" json:"log"`
	DBConfig        *DBConfig     `yaml:"db" json:"db"`
	CacheConfig     *CacheConfig  `yaml:"cache" json:"cache"`
}

//LogConfig 日志配置
type LogConfig struct {
	Level         string `json:"level" yaml:"level"`                 //日志级别
	EnableConsole bool   `yaml:"enableConsole" json:"enableConsole"` //日志是否输出到终端
	Rotated       bool   `json:"rotated" yaml:"rotated"`             //日志是否被分割
	FileDir       string `json:"fileDir" yaml:"fileDir"`             //日志文件所在目录
	MaxSize       int    `json:"maxSize" yaml:"maxSize"`             //每个日志文件长度的最大大小，默认100M
	MaxAge        int    `json:"maxAge" yaml:"maxAge"`               //日志保留的最大天数(只保留最近多少天的日志)
	MaxBackups    int    `json:"maxBackups" yaml:"maxBackups"`       //只保留最近多少个日志文件，用于控制程序总日志的大小
	LocalTime     bool   `json:"localtime" yaml:"localtime"`         //是否使用本地时间，默认使用UTC时间
	Compress      bool   `json:"compress" yaml:"compress"`           //是否压缩日志文件，压缩方法gzip
}

//DBConfig 数据库配置
type DBConfig struct {
	Host             string           `yaml:"host" json:"host"`                         //DB IP
	Port             int              `yaml:"port" json:"port"`                         //DB 端口
	User             string           `yaml:"user" json:"user"`                         //用户名
	Password         string           `yaml:"password" json:"password"`                 //密码
	DBType           string           `yaml:"dbType" json:"dbType"`                     //数据库类型
	Database         string           `yaml:"database" json:"database"`                 //数据库名
	Location         string           `yaml:"location" json:"location"`                 //时区
	TablePrefix      string           `yaml:"tablePrefix" json:"tablePrefix"`           //表前缀
	SingularTable    bool             `yaml:"singularTable" json:"singularTable"`       //表复数禁用
	CreateBatchSize  int              `yaml:"createBatchSize" json:"createBatchSize"`   //批量创建数量
	EnableRawSQL     bool             `yaml:"enableRawSql" json:"enableRawSql"`         //打印原生SQL
	MaxFreeConnCount int              `yaml:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int              `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	FreeMaxLifetime  time.Duration    `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   //闲置连接最大存活时间, 单位: 分钟
	TimeFunc         func() time.Time //设置当前时间函数
}

//Dsn 连接URL
func (c *DBConfig) Dsn() string {
	uri := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s`,
		c.User, c.Password, c.Host, c.Port, c.Database, c.Location)
	return uri
}

//CacheConfig 缓存配置
type CacheConfig struct {
	Host             string        `yaml:"host" json:"host"`                         //Redis IP
	Port             int           `yaml:"port" json:"port"`                         //Redis 端口
	Pwd              string        `yaml:"pwd" json:"pwd"`                           //密码
	MaxFreeConnCount int           `json:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数
	MaxOpenConnCount int           `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数
	FreeMaxLifetime  time.Duration `json:"freeMaxLifetime" yaml:"freeMaxLifetime"`   //闲置连接存活的最大时间, 单位: 分钟
	Database         int           `yaml:"database" json:"database"`                 //数据库
	ConnectTimeout   time.Duration `yaml:"connectTimeout" json:"connectTimeout"`     //连接Redis超时时间, 单位: 秒
	ReadTimeout      time.Duration `yaml:"readTimeout" json:"readTimeout"`           //读取超时时间, 单位: 秒
	WriteTimeout     time.Duration `yaml:"writeTimeout" json:"writeTimeout"`         //写超时时间, 单位: 秒
	DefaultTimeOut   int           `yaml:"defaultTimeOut" json:"defaultTimeOut"`     //默认缓存时间, 单位: 秒
}
