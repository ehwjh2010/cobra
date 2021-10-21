package conf

import (
	"fmt"
	"ginLearn/utils"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Conf Config

type Config struct {
	Application string `yaml:"application" json:"application"`
	Debug       bool   `yaml:"debug" json:"debug"`
	Logfile     string `yaml:"logfile" json:"logfile"`
	Mysql       mysql  `yaml:"mysql" json:"mysql"`
	Redis       redis  `yaml:"redis" json:"redis"`
}
type mysql struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}
type redis struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
	Pwd  string `yaml:"pwd" json:"pwd"`
}

func InitConfig() {
	env := os.Getenv("env")

	if utils.IsEmptyStr(env) {
		env = "dev"
	}

	log.Println("Start load config")

	dir, _ := os.Getwd()

	configFileName := fmt.Sprintf("config_%s.yaml", env)

	configFilePath := utils.PathJoin(dir, "conf", configFileName)

	exist, err := utils.EnsurePathExist(configFilePath)

	if err != nil {
		log.Fatalf("Stat config file failed!, %v", err)
	} else if !exist {
		log.Fatalf("Confif file not exist, path is %s", configFilePath)
	}

	yamlFile, err := utils.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Yamlfile.get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatalf("Load config failed! reason: %v", err)
	}

	log.Println("Load config success")
}
