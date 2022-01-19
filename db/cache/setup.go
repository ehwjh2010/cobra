package cache

import (
	"github.com/ehwjh2010/viper/client/enum"
	"github.com/ehwjh2010/viper/client/setting"
	"github.com/ehwjh2010/viper/log"
)

//SetUp 初始化缓存
func SetUp(conf *setting.Cache) (client *RedisClient, err error) {
	c, err := InitCacheWithGoRedis(conf)
	if err != nil {
		log.Err("Connect redis failed", err)
		return nil, err
	}

	log.Debug("Connect redis success")

	if conf.DefaultTimeOut <= 0 {
		conf.DefaultTimeOut = enum.FiveMinute
	}

	client = NewRedisClient(c, conf, conf.DefaultTimeOut)
	client.WatchHeartbeat()

	return client, nil
}
