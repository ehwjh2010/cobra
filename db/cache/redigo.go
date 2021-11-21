package cache

import (
	"fmt"
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/log"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

const network = "tcp"

//InitCacheWithRedisGo 使用RedisGo初始化
func InitCacheWithRedisGo(redisConfig *client.Cache) (*redis.Pool, error) {

	if redisConfig == nil {
		return nil, nil
	}

	// 建立连接池
	redisClient := &redis.Pool{
		MaxIdle:     redisConfig.MaxFreeConnCount,                             // 最大空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
		MaxActive:   redisConfig.MaxOpenConnCount,                             // 最大连接数，表示同时最多有N个连接 ，为0表示没有限制
		IdleTimeout: time.Duration(redisConfig.FreeMaxLifetime) * time.Minute, // 最大空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Wait:        true,                                                     // 当链接数达到最大后是否阻塞，如果不的话，达到最大后返回错误
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(network, fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
				redis.DialPassword(redisConfig.Pwd),
				redis.DialDatabase(redisConfig.Database),
				redis.DialConnectTimeout(time.Duration(redisConfig.ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(redisConfig.ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(redisConfig.WriteTimeout)*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	log.Debug("Connect redis success!")

	return redisClient, nil
}

func CloseCacheWithRedisGo(redisClient *redis.Pool) error {

	if redisClient == nil {
		return nil
	}

	err := redisClient.Close()
	if err != nil {
		log.Error("Close redis failed", zap.String("err", err.Error()))
	} else {
		log.Debug("Close redis success!")
	}

	return err
}
