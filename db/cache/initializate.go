package cache

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/config"
)

//InitCache 初始化缓存
func InitCache(conf *client.Cache) (client *RedisClient, err error) {
	pool, err := InitCacheWithRedisGo(conf)
	if err != nil {
		return nil, err
	}

	if conf.DefaultTimeOut <= 0 {
		conf.DefaultTimeOut = config.FiveMinute
	}

	client = &RedisClient{
		pool:           pool,
		defaultTimeOut: conf.DefaultTimeOut,
	}

	return client, err
}

//Close 关闭连接池
func (c *RedisClient) Close() error {
	return CloseCacheWithRedisGo(c.pool)
}
