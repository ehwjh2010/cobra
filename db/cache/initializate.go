package cache

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/config"
)

//InitCache 初始化缓存
func InitCache(conf *client.Cache) (client *RedisClient, err error) {
	c, err := InitCacheWithGoRedis(conf)
	if err != nil {
		return nil, err
	}

	if conf.DefaultTimeOut <= 0 {
		conf.DefaultTimeOut = config.FiveMinute
	}

	client = NewRedisClient(c, conf.DefaultTimeOut)

	return client, nil
}
