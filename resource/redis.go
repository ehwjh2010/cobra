package resource

import (
	"ginLearn/utils"
)

var RedisClient *utils.RedisClient

func LoadRedis() {
	client, err := utils.InitCache(Conf.RedisConfig)
	if err != nil {
		Logger.PanicF("Load redis failed!, err: %v\n", err)
	}

	RedisClient = client

}
