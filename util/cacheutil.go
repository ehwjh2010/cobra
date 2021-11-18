package util

import (
	"encoding/json"
	"errors"
	myRedis "ginLearn/db/redis"
	"ginLearn/log"
	"ginLearn/types"
	"ginLearn/util/jsonutils"
	"github.com/gomodule/redigo/redis"
	"time"
)

const DefaultTimeOut = 60 * 5 //5分钟

//CacheConfig 缓存配置
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
	DefaultTimeOut   int           `yaml:"defaultTimeOut" json:"defaultTimeOut"`     //默认缓存时间, 单位: 秒
}

func NewCacheConfig() *CacheConfig {
	return &CacheConfig{}
}

type RedisConfigOption func(*CacheConfig)

type RedisClient struct {
	//pool redis连接池
	pool *redis.Pool

	//defaultTimeOut 默认超时时间
	defaultTimeOut int
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

func RedisClientWithDefaultTimeOut(defaultTimeOut int) RedisClientOption {
	return func(client *RedisClient) {
		timeout := defaultTimeOut
		if timeout <= 0 {
			timeout = DefaultTimeOut
		}

		client.defaultTimeOut = timeout
	}
}

//InitCache 初始化缓存
func InitCache(config *CacheConfig) (client *RedisClient, err error) {
	pool, err := myRedis.InitCacheWithRedisGo(config)
	if err != nil {
		return nil, err
	}

	client = NewRedisClient(RedisClientWithPool(pool), RedisClientWithDefaultTimeOut(config.DefaultTimeOut))
	return client, err
}

//Close 关闭连接池
func (c *RedisClient) Close() error {
	return myRedis.CloseCacheWithRedisGo(c.pool)
}

//Set 如果ex小于0, 则使用默认超时时间, ex 单位: 秒
func (c *RedisClient) Set(key string, value interface{}, timeout int) error {
	if value == nil {
		return nil
	}

	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	var err error

	if timeout > 0 {
		_, err = conn.Do("SET", key, value, "EX", timeout)
	} else {
		_, err = conn.Do("SET", key, value, "EX", c.defaultTimeOut)
	}

	if err != nil {
		log.Error("set error", err.Error())
		return err
	}
	return nil
}

//SetJson 设置json
func (c *RedisClient) SetJson(key string, data interface{}, timeout int) error {
	if data == nil {
		return nil
	}

	value, _ := jsonutils.Marshal(data)
	return c.Set(key, value, timeout)
}

//GetString redis get
func (c *RedisClient) GetString(key string) (types.NullString, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewStrNull(), nil
		} else {
			return types.NewStrNull(), err
		}
	}

	return types.NewStr(val), nil
}

//GetBool redis get
func (c *RedisClient) GetBool(key string) (types.NullBool, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Bool(conn.Do("GET", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewBoolNull(), nil
		} else {
			return types.NewBoolNull(), err
		}
	}

	return types.NewBool(val), nil
}

//SetExp 设置过期时间
func (c *RedisClient) SetExp(key string, ex int) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	_, err := conn.Do("EXPIRE", key, ex)
	if err != nil {
		log.Errorf("set error, %v", err.Error())
		return err
	}
	return nil
}

//GetJson 获取JSON
func (c *RedisClient) GetJson(key string, data interface{}) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	bv, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Error("get json error", err.Error())
		return err
	}
	errJson := jsonutils.Unmarshal(bv, data)
	if errJson != nil {
		log.Error("json nil", err.Error())
		return err
	}
	return nil
}

//HSet 对应hset命令
func (c *RedisClient) HSet(key string, field string, data interface{}) error {
	if data == nil {
		return nil
	}

	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	_, err := conn.Do("HSET", key, field, data)
	if err != nil {
		log.Error("hSet error", err.Error())
		return err
	}
	return nil
}

//HGetStr 对应hget命令
func (c *RedisClient) HGetStr(key, field string) (types.NullString, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewStrNull(), nil
		} else {
			return types.NewStrNull(), err
		}
	}
	return types.NewStr(val), nil
}

//HGetInt 对应hget命令
func (c *RedisClient) HGetInt(key, field string) (types.NullInt, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Int(conn.Do("HGET", key, field))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewIntNull(), nil
		} else {
			return types.NewIntNull(), err
		}
	}

	return types.NewInt(val), nil
}

//HGetInt64 对应hget命令
func (c *RedisClient) HGetInt64(key, field string) (types.NullInt64, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Int64(conn.Do("HGET", key, field))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewInt64Null(), nil
		} else {
			return types.NewInt64Null(), err
		}
	}

	return types.NewInt64(val), nil
}

//HGetBool 对应hget命令
func (c *RedisClient) HGetBool(key, field string) (types.NullBool, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Bool(conn.Do("HGET", key, field))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewBoolNull(), nil
		} else {
			return types.NewBoolNull(), err
		}
	}
	return types.NewBool(val), nil
}

//HGetAll 对应hgetall命令
func (c *RedisClient) HGetAll(key string) (map[string]interface{}, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	var val map[string]interface{}

	tmp, err := redis.Bytes(conn.Do("HGETALL", key))
	if err != nil {
		log.Error("hGetAll error", err.Error())
		return nil, err
	}

	err = json.Unmarshal(tmp, &val)
	if err != nil {
		log.Error("json nil, ", err.Error())
		return nil, err
	}
	return val, nil
}

//Incr 对应incr命令
func (c *RedisClient) Incr(key string) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		log.Error("INCR error, ", err.Error())
		return 0, err
	}
	return val, nil

}

//IncrBy 对应incrby命令
func (c *RedisClient) IncrBy(key string, n int) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Int(conn.Do("INCRBY", key, n))
	if err != nil {
		log.Error("INCRBY error", err.Error())
		return 0, err
	}
	return val, nil
}

//Decr 对应decr命令
func (c *RedisClient) Decr(key string) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Int(conn.Do("DECR", key))
	if err != nil {
		log.Error("DECR error", err.Error())
		return 0, err
	}
	return val, nil
}

//DecrBy 对应decrby命令
func (c *RedisClient) DecrBy(key string, n int) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Int(conn.Do("DECRBY", key, n))
	if err != nil {
		log.Error("DECRBY error", err.Error())
		return 0, err
	}
	return val, nil
}

//SAdd 对应sadd命令
func (c *RedisClient) SAdd(key string, v interface{}) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	_, err := conn.Do("SADD", key, v)
	if err != nil {
		log.Error("SADD error", err.Error())
		return err
	}
	return nil
}

//SMembersStr 对应smembers命令
func (c *RedisClient) SMembersStr(key string) ([]string, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Error("json nil", err)
		return nil, err
	}
	return val, nil
}

//SMembersInt 对应smembers命令
func (c *RedisClient) SMembersInt(key string) ([]int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Ints(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Error("json nil", err)
		return nil, err
	}
	return val, nil
}

//SMembersInt64 对应smembers命令
func (c *RedisClient) SMembersInt64(key string) ([]int64, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Int64s(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Error("json nil", err)
		return nil, err
	}
	return val, nil
}

//SISMembers 对应sismembers命令
func (c *RedisClient) SISMembers(key, v string) bool {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Bool(conn.Do("SISMEMBER", key, v))
	if err != nil {
		log.Error("SISMEMBER error", err.Error())
		return false
	}
	return val
}

//Exist 对应exists命令
func (c *RedisClient) Exist(key string) bool {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	val, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		log.Error(err)
		return false
	}
	return val
}

//Del 对应del命令
func (c *RedisClient) Del(key string) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("Redis conn close failed, %v", err)
		}
	}()

	_, err := conn.Do("DEL", key)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
