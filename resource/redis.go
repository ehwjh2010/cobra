package resource

import (
	"ginLearn/utils"
	"log"
)

var RedisClient utils.RedisClient

func LoadCache() {
	err := utils.InitCache(Conf.RedisConfig, &RedisClient)
	if err != nil {
		log.Panicf("Load redis failed!, err: %v\n", err)
	}

}
