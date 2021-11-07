package utils

import (
	"encoding/json"
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

func NewCacheConfig() *CacheConfig {
	return &CacheConfig{}
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

func InitCache(config *CacheConfig) (client *RedisClient, err error) {
	pool, err := InitCacheWithRedisGo(config)
	if err != nil {
		return nil, err
	}

	client = NewRedisClient(RedisClientWithPool(pool))
	return client, err
}

func (c *RedisClient) Close() error {
	return CloseCacheWithRedisGo(c.pool)
}

//Set 如果ex小于0, 则认为没有设置时间, ex 单位: 秒
func (c *RedisClient) Set(key string, v interface{}, timeout int) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	var err error

	if timeout > 0 {
		_, err = conn.Do("SET", key, v, "EX", timeout)
	} else {
		_, err = conn.Do("SET", key, v)
	}

	if err != nil {
		Error("set error", err.Error())
		return err
	}
	return nil
}

//SetJson 设置json
func (c *RedisClient) SetJson(key string, data interface{}, timeout int) error {
	value, _ := JsonMarshal(data)
	return c.Set(key, value, timeout)
}

//GetString redis get
func (c *RedisClient) GetString(key string) (string, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		Error("get error", err.Error())
		return "", err
	}

	return val, nil
}

//GetBool redis get
func (c *RedisClient) GetBool(key string) (bool, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	val, err := redis.Bool(conn.Do("GET", key))
	if err != nil {
		Error("get error", err.Error())
		return false, err
	}

	return val, nil
}

func (c *RedisClient) SetExp(key string, ex int) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	_, err := conn.Do("EXPIRE", key, ex)
	if err != nil {
		Error("set error", err.Error())
		return err
	}
	return nil
}

func (c *RedisClient) GetJson(key string, data interface{}) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	bv, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		Error("get json error", err.Error())
		return err
	}
	errJson := JsonUnmarshal(bv, data)
	if errJson != nil {
		Error("json nil", err.Error())
		return err
	}
	return nil
}

func (c *RedisClient) HSet(key string, field string, data interface{}) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	_, err := conn.Do("HSET", key, field, data)
	if err != nil {
		Error("hSet error", err.Error())
		return err
	}
	return nil
}

func (c *RedisClient) HGetStr(key, field string) (string, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	data, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		Error("hGet error", err.Error())
		return "", err
	}
	return data, nil
}

func (c *RedisClient) HGetInt(key, field string) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	data, err := redis.Int(conn.Do("HGET", key, field))
	if err != nil {
		Error("hGet error", err.Error())
		return 0, err
	}
	return data, nil
}

func (c *RedisClient) HGetInt64(key, field string) (int64, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	data, err := redis.Int64(conn.Do("HGET", key, field))
	if err != nil {
		Error("hGet error", err.Error())
		return 0, err
	}
	return data, nil
}

func (c *RedisClient) HGetBool(key, field string) (bool, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	data, err := redis.Bool(conn.Do("HGET", key, field))
	if err != nil {
		Error("hGet error", err.Error())
		return false, err
	}
	return data, nil
}

func (c *RedisClient) HGetAll(key string) (map[string]interface{}, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	var data map[string]interface{}

	tmp, err := redis.Bytes(conn.Do("HGETALL", key))
	if err != nil {
		Error("hGetAll error", err.Error())
		return nil, err
	}

	err = json.Unmarshal(tmp, &data)
	if err != nil {
		Error("json nil, ", err.Error())
		return nil, err
	}
	return data, nil
}

func (c *RedisClient) Incr(key string) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	count, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		Error("INCR error", err.Error())
		return 0, err
	}
	return count, nil

}

func (c *RedisClient) IncrBy(key string, n int) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	count, err := redis.Int(conn.Do("INCRBY", key, n))
	if err != nil {
		Error("INCRBY error", err.Error())
		return 0, err
	}
	return count, nil
}

func (c *RedisClient) Decr(key string) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	count, err := redis.Int(conn.Do("DECR", key))
	if err != nil {
		Error("DECR error", err.Error())
		return 0, err
	}
	return count, nil
}

func (c *RedisClient) DecrBy(key string, n int) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	count, err := redis.Int(conn.Do("DECRBY", key, n))
	if err != nil {
		Error("DECRBY error", err.Error())
		return 0, err
	}
	return count, nil
}

func (c *RedisClient) SAdd(key string, v interface{}) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	_, err := conn.Do("SADD", key, v)
	if err != nil {
		Error("SADD error", err.Error())
		return err
	}
	return nil
}

func (c *RedisClient) SMembersStr(key string) ([]string, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	data, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		Error("json nil", err)
		return nil, err
	}
	return data, nil
}

func (c *RedisClient) SMembersInt(key string) ([]int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	data, err := redis.Ints(conn.Do("SMEMBERS", key))
	if err != nil {
		Error("json nil", err)
		return nil, err
	}
	return data, nil
}

func (c *RedisClient) SMembersInt64(key string) ([]int64, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	data, err := redis.Int64s(conn.Do("SMEMBERS", key))
	if err != nil {
		Error("json nil", err)
		return nil, err
	}
	return data, nil
}

func (c *RedisClient) SISMembers(key, v string) bool {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	b, err := redis.Bool(conn.Do("SISMEMBER", key, v))
	if err != nil {
		Error("SISMEMBER error", err.Error())
		return false
	}
	return b
}

func (c *RedisClient) Exist(key string) bool {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	b, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		Error(err)
		return false
	}
	return b
}

func (c *RedisClient) Del(key string) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			Errorf("Redis conn close failed, err: %v", err)
		}
	}()

	_, err := conn.Do("DEL", key)
	if err != nil {
		Error(err)
		return err
	}
	return nil
}
