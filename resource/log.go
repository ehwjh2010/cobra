package resource

import (
	"ginLearn/utils"
	"log"
)

func LoadLog() {

	err := utils.InitLogrus(Conf.Application, Conf.LogConfig)

	if err != nil {
		log.Panicf("Load config failed!, err: %v", err)
	}

	log.Println("Load log success")
}

func ReleaseLog() error {
	return utils.CloseLogrus()
}
