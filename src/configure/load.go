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

	currentDir, _ := os.Getwd()

	localConfigPath := utils.PathJoin(currentDir, "conf", "config_local.yaml")

	exist, err := utils.EnsurePathExist(localConfigPath)

	if utils.IsNotNil(err) {
		log.Fatalf("State local config failed! err: %s", err)
	}

	var configFilePath string

	if exist {
		configFilePath = localConfigPath
	} else {
		env := os.Getenv("ENV")

		if utils.IsEmptyStr(env) {
			env = "dev"
		}

		configFileName := fmt.Sprintf("config_%s.yaml", env)

		configFilePath = utils.PathJoin(currentDir, "conf", configFileName)

		exist, err = utils.EnsurePathExist(configFilePath)

		if utils.IsNotNil(err) {
			log.Fatalf("State %s config failed! err: %s", env, err)
		} else if !exist {
			log.Fatalf("Confif file not exist, path is %s", configFilePath)
		}
	}

	yamlFile, err := utils.ReadFile(configFilePath)
	if utils.IsNotNil(err) {
		log.Fatalf("Yamlfile.get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if utils.IsNotNil(err) {
		log.Fatalf("Load config failed! reason: %v", err)
	}

	log.Println("Load config success")
}
