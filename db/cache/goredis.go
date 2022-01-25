package cache

import (
	"context"
	"fmt"
	"github.com/ehwjh2010/viper/client/settings"
	"github.com/go-redis/redis/v8"
	"time"
)

const network = "tcp"

func InitCacheWithGoRedis(redisConfig settings.Cache) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Network:      network,
		Addr:         fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Username:     redisConfig.User,
		Password:     redisConfig.Pwd,
		DB:           redisConfig.Database,
		DialTimeout:  time.Duration(redisConfig.ConnectTimeout) * time.Second,
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(redisConfig.WriteTimeout) * time.Second,
		PoolFIFO:     true,
		PoolSize:     redisConfig.MaxOpenConnCount,
		MinIdleConns: redisConfig.MinFreeConnCount,
		MaxRetries:   redisConfig.MaxRetries,
		IdleTimeout:  time.Duration(redisConfig.FreeMaxLifetime) * time.Second,
		MaxConnAge:   time.Duration(redisConfig.ConnMaxLifetime) * time.Second,
	})

	ctx := context.TODO()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
