package utils

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

const network = "tcp"

func InitCacheWithRedisGo(redisConfig *CacheConfig) (*redis.Pool, error) {

	if redisConfig == nil {
		return nil, nil
	}

	// 建立连接池
	redisClient := &redis.Pool{
		MaxIdle:     redisConfig.MaxFreeConnCount,
		MaxActive:   redisConfig.MaxOpenConnCount,
		IdleTimeout: redisConfig.FreeMaxLifetime * time.Minute,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(network, fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
				redis.DialPassword(redisConfig.Pwd),
				redis.DialDatabase(redisConfig.Database),
				redis.DialConnectTimeout(redisConfig.ConnectTimeout*time.Second), //默认时间: 30s
				redis.DialReadTimeout(redisConfig.ReadTimeout*time.Second),
				redis.DialWriteTimeout(redisConfig.WriteTimeout*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

	log.Println("Connect redis success!")

	return redisClient, nil
}

func CloseCacheWithRedisGo(redisClient *redis.Pool) error {

	if redisClient == nil {
		return nil
	}

	err := redisClient.Close()
	if err != nil {
		log.Printf("Close redis failed, err: %v", err)
	} else {
		log.Println("Close redis success!")
	}

	return err
}
