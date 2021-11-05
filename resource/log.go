package resource

import (
	"ginLearn/utils"
	"log"
)

var Logger *utils.Logrus

func LoadLog() {
	Logger = utils.NewLogrus(Conf.LogConfig)

	err := Logger.SetUp(Conf.Application)

	if err != nil {
		log.Panicf("Load config failed!, err: %v", err)
	}

	log.Println("Load log success")
}
