package cache

import (
	"github.com/ehwjh2010/cobra/client"
)

//InitCache 初始化缓存
func InitCache(config *client.CacheConfig) (client *RedisClient, err error) {
	pool, err := InitCacheWithRedisGo(config)
	if err != nil {
		return nil, err
	}

	client = NewRedisClient(
		RedisClientWithPool(pool),
		RedisClientWithDefaultTimeOut(config.DefaultTimeOut))
	return client, err
}

//Close 关闭连接池
func (c *RedisClient) Close() error {
	return CloseCacheWithRedisGo(c.pool)
}
