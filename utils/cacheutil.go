package utils

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type CacheConfig struct {
	Host             string        `yaml:"host" json:"host"`                         //Redis IP
	Port             int           `yaml:"port" json:"port"`                         //Redis 端口
	Pwd              string        `yaml:"pwd" json:"pwd"`                           //密码
	MaxFreeConnCount int           `json:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int           `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	FreeMaxLifetime  time.Duration `json:"freeMaxLifetime" yaml:"freeMaxLifetime"`   //闲置连接存活的最大时间, 单位: 分钟
	Database         int           `yaml:"database" json:"database"`                 //数据库
	ConnectTimeout   time.Duration `yaml:"connectTimeout" json:"connectTimeout"`     //连接Redis超时时间, 单位: 秒
	ReadTimeout      time.Duration `yaml:"readTimeout" json:"readTimeout"`           //读取超时时间, 单位: 秒
	WriteTimeout     time.Duration `yaml:"writeTimeout" json:"writeTimeout"`         //写超时时间, 单位: 秒
}

type RedisConfigOption func(*CacheConfig)

type RedisClient struct {
	pool *redis.Pool
}

func NewRedisClient(args ...RedisClientOption) (client *RedisClient) {
	client = &RedisClient{}
	for _, arg := range args {
		arg(client)
	}
	return client
}

type RedisClientOption func(client *RedisClient)

func RedisClientWithPool(pool *redis.Pool) RedisClientOption {
	return func(client *RedisClient) {
		client.pool = pool
	}
}

func RedisConfigWithHost(host string) RedisConfigOption {
	return func(config *CacheConfig) {
		config.Host = host
	}
}

func RedisConfigWithPort(port int) RedisConfigOption {
	return func(config *CacheConfig) {
		config.Port = port
	}
}

func RedisConfigWithMaxFreeConnCount(maxFreeConnCount int) RedisConfigOption {
	return func(config *CacheConfig) {
		config.MaxFreeConnCount = maxFreeConnCount
	}
}

func RedisConfigWithMaxOpenConnCount(maxOpenConnCount int) RedisConfigOption {
	return func(config *CacheConfig) {
		config.MaxOpenConnCount = maxOpenConnCount
	}
}

func RedisConfigWithFreeMaxLifetime(freeMaxLifetime time.Duration) RedisConfigOption {
	return func(config *CacheConfig) {
		config.FreeMaxLifetime = freeMaxLifetime
	}
}

func RedisConfigWithDatabase(database int) RedisConfigOption {
	return func(config *CacheConfig) {
		config.Database = database
	}
}

func RedisConfigWithConnectTimeout(connectTimeout time.Duration) RedisConfigOption {
	return func(config *CacheConfig) {
		config.ConnectTimeout = connectTimeout
	}
}

func RedisConfigWithReadTimeout(readTimeout time.Duration) RedisConfigOption {
	return func(config *CacheConfig) {
		config.ReadTimeout = readTimeout
	}
}

func RedisConfigWithWriteTimeout(writeTimeout time.Duration) RedisConfigOption {
	return func(config *CacheConfig) {
		config.WriteTimeout = writeTimeout
	}
}

func InitCache(config *CacheConfig) (*RedisClient, error) {
	pool, err := InitCacheWithRedisGo(config)
	if err != nil {
		return nil, err
	}

	client := NewRedisClient(RedisClientWithPool(pool))
	return client, err
}

func (c *RedisClient) Close() error {
	return CloseCacheWithRedisGo(c.pool)
}
