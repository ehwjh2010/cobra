package resource

import (
	"ginLearn/utils"
	"log"
)

func LoadLogrus() {

	err := utils.InitLogrus(Conf.Application, Conf.LogConfig)

	if err != nil {
		log.Panicf("Load config failed!, err: %v", err)
	}
}

func ReleaseLogrus() error {
	return utils.CloseLogFile()
}
