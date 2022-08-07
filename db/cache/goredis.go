package cache

import (
	"context"
	"time"

	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/go-redis/redis/v8"
)

const (
	network          = "tcp"
	defaultAddr      = "localhost:6379"
	connectTimeout   = enums.FiveSecond
	readTimeout      = enums.ThreeSecond
	waitTimeOut      = enums.ThreeSecond
	freeMaxLifetime  = enums.TenMinute
	connMaxLifetime  = enums.OneHour
	noTimeOut        = -1
	maxOpenConnCount = 10
	minFreeConnCount = 3
	maxRetries       = 3
)

func initCacheWithGoRedis(conf *Cache) (*redis.Client, error) {
	// 设置默认配置
	setRedisDefaultConf(conf)

	redisClient := redis.NewClient(&redis.Options{
		// 连接信息
		Network:  network,
		Addr:     conf.Addr,
		Username: conf.User,
		Password: conf.Pwd,
		DB:       conf.Database,

		// 超时
		DialTimeout:  integerToTime(conf.ConnectTimeout),
		ReadTimeout:  integerToTime(conf.ReadTimeout),
		WriteTimeout: integerToTime(conf.WriteTimeout),
		PoolTimeout:  integerToTime(conf.BusyWaitTimeOut),
		MaxConnAge:   integerToTime(conf.ConnMaxLifetime),

		// 连接配置
		PoolFIFO:     true,
		PoolSize:     conf.MaxOpenConnCount,
		MinIdleConns: conf.MinFreeConnCount,
		MaxRetries:   conf.MaxRetries,
		IdleTimeout:  integerToTime(conf.FreeMaxLifetime),
	})

	ctx := context.TODO()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}

func integerToTime(ts int) time.Duration {
	var result time.Duration
	switch ts {
	case noTimeOut:
		result = time.Duration(noTimeOut)
	default:
		result = time.Duration(ts) * time.Second
	}
	return result
}

func setRedisDefaultConf(conf *Cache) {
	if str.IsEmpty(conf.Network) {
		conf.Network = network
	}

	if str.IsEmpty(conf.Addr) {
		conf.Addr = defaultAddr
	}

	if conf.ConnectTimeout <= 0 {
		conf.ConnectTimeout = connectTimeout
	}

	if conf.ReadTimeout < 0 {
		conf.ReadTimeout = noTimeOut
	}

	if conf.ReadTimeout == 0 {
		conf.ReadTimeout = readTimeout
	}

	if conf.WriteTimeout < 0 {
		conf.WriteTimeout = noTimeOut
	}

	if conf.WriteTimeout == 0 {
		conf.WriteTimeout = conf.ReadTimeout
	}

	if conf.BusyWaitTimeOut <= 0 {
		conf.BusyWaitTimeOut = waitTimeOut
	}

	if conf.MaxOpenConnCount <= 0 {
		conf.MaxOpenConnCount = maxOpenConnCount
	}

	if conf.MinFreeConnCount <= 0 {
		conf.MinFreeConnCount = minFreeConnCount
	}

	if conf.MaxRetries <= 0 {
		conf.MaxRetries = maxRetries
	}

	if conf.FreeMaxLifetime < 0 {
		conf.FreeMaxLifetime = -1
	}

	if conf.FreeMaxLifetime == 0 {
		conf.FreeMaxLifetime = freeMaxLifetime
	}

	if conf.ConnMaxLifetime <= 0 {
		conf.ConnMaxLifetime = connMaxLifetime
	}

	if conf.PeriodMillSec <= 0 {
		conf.PeriodMillSec = 3000
	}

	if conf.AddMillSec <= 0 {
		conf.AddMillSec = 5000
	}
}
