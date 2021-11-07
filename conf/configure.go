package conf

import (
	"fmt"
	"ginLearn/utils"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
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

type LogConfig struct {
	//Level 日志级别
	Level string `json:"level" yaml:"level"`

	//EnableConsole 日志是否输出到终端
	EnableConsole bool `yaml:"enableConsole" json:"enableConsole"`

	//Rotated 日志是否被分割
	Rotated bool `json:"rotated" yaml:"rotated"`

	//FileDir 日志文件所在目录
	FileDir string `json:"fileDir" yaml:"fileDir"`

	//MaxSize 每个日志文件长度的最大大小，默认100M
	MaxSize int `json:"maxSize" yaml:"maxSize"`

	//MaxAge 日志保留的最大天数(只保留最近多少天的日志)
	MaxAge int `json:"maxAge" yaml:"maxAge"`

	//MaxBackups 只保留最近多少个日志文件，用于控制程序总日志的大小
	MaxBackups int `json:"maxBackups" yaml:"maxBackups"`

	//LocalTime 是否使用本地时间，默认使用UTC时间
	LocalTime bool `json:"localtime" yaml:"localtime"`

	//Compress 是否压缩日志文件，压缩方法gzip
	Compress bool `json:"compress" yaml:"compress"`
}

type DBConfig struct {
	Host             string        `yaml:"host" json:"host"`                         //DB IP
	Port             int           `yaml:"port" json:"port"`                         //DB 端口
	User             string        `yaml:"user" json:"user"`                         //用户名
	Password         string        `yaml:"password" json:"password"`                 //密码
	DBType           string        `yaml:"dbType" json:"dbType"`                     //数据库类型
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

type CacheConfig struct {
	Host             string        `yaml:"host" json:"host"`                         //Redis IP
	Port             int           `yaml:"port" json:"port"`                         //Redis 端口
	Pwd              string        `yaml:"pwd" json:"pwd"`                           //密码
	MaxFreeConnCount int           `json:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int           `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	FreeMaxLifetime  time.Duration `json:"freeMaxLifetime" yaml:"freeMaxLifetime"`   //闲置连接存活的最大时间, 单位: 分钟
	Database         int           `yaml:"database" json:"database"`                 //数据库
	ConnectTimeout   time.Duration `yaml:"connectTimeout" json:"connectTimeout"`     //连接Redis超时时间, 单位: 秒
	ReadTimeout      time.Duration `yaml:"readTimeout" json:"readTimeout"`           //读取超时时间, 单位: 秒
	WriteTimeout     time.Duration `yaml:"writeTimeout" json:"writeTimeout"`         //写超时时间, 单位: 秒
}

var Conf Config

//LoadConfig 从配置文件中加载配置
func LoadConfig() {

	log.Println("Start load config")

	configFilePath, err := ensureConfigPath()

	if err != nil {
		panic(err)
	}

	yamlFile, err := utils.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		panic(err)
	}

	//设置环境标识
	env := getEnv()

	Conf.Env = env

	log.Println("Load config success")
}

//ensureConfigPath 确定配置文件
func ensureConfigPath() (string, error) {
	currentDir, _ := os.Getwd()

	//优先读取本地配置, 利于本地开发以及线上配置
	localConfigPath := utils.PathJoin(currentDir, "conf", "config_local.yaml")

	exist, err := utils.EnsurePathExist(localConfigPath)

	if err != nil {
		return "", err
	}

	if exist {
		return localConfigPath, nil
	}

	//未读取到local配置文件, 则读取相应环境配置文件
	env := getEnv()

	configFileName := fmt.Sprintf("config_%s.yaml", strings.ToLower(env))

	configFilePath := utils.PathJoin(currentDir, "conf", configFileName)

	exist, err = utils.EnsurePathExist(configFilePath)

	if err != nil {
		return "", err
	} else if !exist {
		return "", err
	}

	return configFilePath, nil
}

//getEnv 获取环境标识
func getEnv() string {
	env := os.Getenv("ENV")

	if utils.IsEmptyStr(env) {
		env = "dev"
	}

	return env
}
