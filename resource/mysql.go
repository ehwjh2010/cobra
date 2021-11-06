package resource

import (
	"ginLearn/utils"
	"log"
)

var DBClient utils.DBClient

func LoadDB() {

	err := utils.InitDB(Conf.DBConfig, &DBClient)

	if err != nil {
		log.Panicf("Load mysql failed!, err: %v", err)
	}
}
