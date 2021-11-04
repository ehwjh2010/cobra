package utils

import (
	"ginLearn/client/setting"
	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	pool *redis.Pool
}

func (c *RedisClient) SetUp(redisConfig *setting.RedisConfig) error {
	pool, err := InitCacheWithRedisGo(redisConfig)
	if err != nil {
		return err
	}

	c.pool = pool
	return nil
}

func (c *RedisClient) Close() error {
	return CloseCacheWithRedisGo(c.pool)
}
