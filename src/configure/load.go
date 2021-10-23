package configure

import (
	"fmt"
	"ginLearn/client/setting"
	"ginLearn/utils"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Conf setting.Config

//LoadConfig 从配置文件中加载配置
func LoadConfig() {

	log.Println("Start load config")

	configFilePath := ensureConfigPath()

	yamlFile, err := utils.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Yamlfile.get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatalf("Load config failed! reason: %v", err)
	}

	log.Println("Load config success")

	log.Printf("%#v", Conf)
}

//ensureConfigPath 确定配置文件
func ensureConfigPath() string {
	currentDir, _ := os.Getwd()

	//优先读取本地配置, 利于本地开发以及线上配置
	localConfigPath := utils.PathJoin(currentDir, "conf", "config_local.yaml")

	exist, err := utils.EnsurePathExist(localConfigPath)

	if err != nil {
		log.Fatalf("State local config failed! err: %s", err)
	}

	if exist {
		return localConfigPath
	}

	//未读取到local配置文件, 则读取相应环境配置文件
	env := os.Getenv("ENV")

	if utils.IsEmptyStr(env) {
		env = "dev"
	}

	configFileName := fmt.Sprintf("config_%s.yaml", env)

	configFilePath := utils.PathJoin(currentDir, "conf", configFileName)

	exist, err = utils.EnsurePathExist(configFilePath)

	if err != nil {
		log.Fatalf("State %s config failed! err: %s", env, err)
	} else if !exist {
		log.Fatalf("%s config file not exist, path is %s", env, configFilePath)
	}

	return configFilePath
}
