package cache

import (
	"github.com/ehwjh2010/cobra/client"
)

//InitCache 初始化缓存
func InitCache(config *client.Cache) (client *RedisClient, err error) {
	pool, err := InitCacheWithRedisGo(config)
	if err != nil {
		return nil, err
	}

	if config.DefaultTimeOut <= 0 {
		config.DefaultTimeOut = DefaultTimeOut
	}

	client = &RedisClient{
		pool:           pool,
		defaultTimeOut: config.DefaultTimeOut,
	}

	return client, err
}

//Close 关闭连接池
func (c *RedisClient) Close() error {
	return CloseCacheWithRedisGo(c.pool)
}
