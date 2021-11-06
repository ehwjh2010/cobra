package resource

import (
	"ginLearn/utils"
	"log"
)

func LoadLog() {
	err := utils.InitLog(Conf.LogConfig, Conf.Application)

	if err != nil {
		log.Panicf("Load log failed!, err: %v", err)
	}

	log.Println("Load log success")
}
