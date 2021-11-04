package resource

import (
	"ginLearn/utils"
)

var RedisClient utils.RedisClient

func LoadRedis() {
	err := RedisClient.SetUp(Conf.RedisConfig)
	if err != nil {
		utils.PanicF("Load redis failed!, err: %v\n", err)
	}

}
