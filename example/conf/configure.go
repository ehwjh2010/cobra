package conf

import (
	"fmt"
	"cobra/client"
	"cobra/util/fileutils"
	"cobra/util/pathutils"
	"cobra/util/strutils"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Env             string              `yaml:"env" json:"env"`                         //环境标识
	Application     string              `yaml:"application" json:"application"`         //应用名
	ServerPort      uint                `yaml:"serverPort" json:"serverPort"`           //服务器端口
	ShutDownTimeout time.Duration       `yaml:"shutDownTimeout" json:"shutDownTimeout"` //优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒
	LogConfig       *client.LogConfig   `yaml:"log" json:"log"`
	DBConfig        *client.DBConfig    `yaml:"db" json:"db"`
	CacheConfig     *client.CacheConfig `yaml:"cache" json:"cache"`
}

var Conf Config

//LoadConfig 从配置文件中加载配置
func LoadConfig() {

	log.Println("Start load config")

	configFilePath, err := ensureConfigPath()

	if err != nil {
		panic(err)
	}

	yamlFile, err := fileutils.ReadFile(configFilePath)
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
	localConfigPath := pathutils.PathJoin(currentDir, "conf", "config_local.yaml")

	exist, err := pathutils.EnsurePathExist(localConfigPath)

	if err != nil {
		return "", err
	}

	if exist {
		return localConfigPath, nil
	}

	//未读取到local配置文件, 则读取相应环境配置文件
	env := getEnv()

	configFileName := fmt.Sprintf("config_%s.yaml", strings.ToLower(env))

	configFilePath := pathutils.PathJoin(currentDir, "conf", configFileName)

	exist, err = pathutils.EnsurePathExist(configFilePath)

	if err != nil {
		return "", err
	} else if !exist {
		return "", err
	}

	return configFilePath, nil
}

//getEnv 获取环境标识
func getEnv() string {
	env := strings.ToLower(os.Getenv("ENV"))

	if strutils.IsEmptyStr(env) {
		env = "dev"
	}

	return env
}
