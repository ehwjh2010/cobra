package cache

import (
	"github.com/ehwjh2010/viper/log"
)

// SetUp 初始化缓存.
func SetUp(conf Cache) (client *RedisClient, err error) {

	c, err := initCacheWithGoRedis(&conf)
	if err != nil {
		log.Errorf("connect redis failed, error: %s", err)
		return nil, err
	}

	log.Debugf("connect redis success")

	client = NewRedisClient(c, &conf)
	client.WatchHeartbeat()

	return client, nil
}
