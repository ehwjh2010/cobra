package cache

import (
	"github.com/ehwjh2010/viper/client/settings"
	"github.com/ehwjh2010/viper/log"
)

// SetUp 初始化缓存
func SetUp(conf *settings.Cache) (client *RedisClient, err error) {

	c, err := initCacheWithGoRedis(conf)
	if err != nil {
		log.Err("Connect redis failed", err)
		return nil, err
	}

	log.Debug("Connect redis success")

	client = NewRedisClient(c, conf)
	client.WatchHeartbeat()

	return client, nil
}
