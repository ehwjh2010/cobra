package cache

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ehwjh2010/cobra/config"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/types"
	"github.com/ehwjh2010/cobra/util/serialize"
	"github.com/ehwjh2010/cobra/util/timer"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

var ErrNullValue = errors.New("dest value is null")

type RedisClient struct {
	//pool redis连接池
	pool *redis.Pool

	//defaultTimeOut 默认过期时间
	defaultTimeOut int
}

//===============================Command Set===================================

//Set 如果timeout小于等于0, 则使用默认超时时间, ex 单位: 秒
func (c *RedisClient) Set(key string, value interface{}, timeout int) error {
	if value == nil {
		return nil
	}

	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	var err error

	if timeout > 0 {
		_, err = conn.Do("SET", key, value, "EX", timeout)
	} else {
		_, err = conn.Do("SET", key, value, "EX", c.defaultTimeOut)
	}

	if err != nil {
		log.Error("Command set", zap.Error(err), zap.String("key", key))
		return err
	}
	return nil
}

//SetTime 设置时间
func (c *RedisClient) SetTime(key string, t time.Time, timeout int) error {
	format := timer.Time2Str(t)

	return c.Set(key, format, timeout)
}

//===============================Command Get===================================

//SetJson 设置json
func (c *RedisClient) SetJson(key string, data interface{}, timeout int) error {
	if data == nil {
		return nil
	}

	value, _ := serialize.Marshal(data)
	return c.Set(key, value, timeout)
}

//GetString redis get
func (c *RedisClient) GetString(key string) (types.NullString, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
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

//GetInt redis get
func (c *RedisClient) GetInt(key string) (types.NullInt, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewIntNull(), nil
		} else {
			return types.NewIntNull(), err
		}
	}

	return types.NewInt(val), nil
}

//GetTime redis get
func (c *RedisClient) GetTime(key string) (types.NullTime, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewTimeNull(), nil
		} else {
			return types.NewTimeNull(), err
		}
	}

	t, err := timer.Str2Time(val)
	if err != nil {
		return types.NewTimeNull(), err
	}

	return types.NewTime(t), nil
}

//GetInt64 redis get
func (c *RedisClient) GetInt64(key string) (types.NullInt64, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Int64(conn.Do("GET", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewInt64Null(), nil
		} else {
			return types.NewInt64Null(), err
		}
	}

	return types.NewInt64(val), nil
}

//GetFloat64 redis get
func (c *RedisClient) GetFloat64(key string) (types.NullFloat64, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Float64(conn.Do("GET", key))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewFloat64Null(), nil
		} else {
			return types.NewFloat64Null(), err
		}
	}

	return types.NewFloat64(val), nil
}

//SetExp 设置过期时间
func (c *RedisClient) SetExp(key string, ex int) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	_, err := conn.Do("EXPIRE", key, ex)
	if err != nil {
		log.Error("Command expire", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	bv, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Error("Command get json", zap.Error(err))
		return err
	}

	if bytes.Equal(bv, config.NullBytes) {
		return ErrNullValue
	}

	errJson := serialize.Unmarshal(bv, data)
	if errJson != nil {
		log.Error("Json unmarshal failed", zap.String("err", errJson.Error()))
		return err
	}
	return nil
}

//===============================Command HSet===================================

//HSet 对应hset命令
func (c *RedisClient) HSet(key string, field string, data interface{}) error {
	if data == nil {
		return nil
	}

	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	_, err := conn.Do("HSET", key, field, data)
	if err != nil {
		log.Error("Command hset", zap.Error(err))
		return err
	}
	return nil
}

//HSetTime 对应hset命令
func (c *RedisClient) HSetTime(key string, field string, t time.Time) error {

	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	data := timer.Time2Str(t)

	_, err := conn.Do("HSET", key, field, data)
	if err != nil {
		log.Error("Command hset", zap.Error(err))
		return err
	}
	return nil
}

//===============================Command HGet===================================

//HGetStr 对应hget命令
func (c *RedisClient) HGetStr(key, field string) (types.NullString, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
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

func (c *RedisClient) HGetTime(key, field string) (types.NullTime, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewTimeNull(), nil
		} else {
			return types.NewTimeNull(), err
		}
	}

	t, err := timer.Str2Time(val)

	if err != nil {
		return types.NewTimeNull(), err
	}

	return types.NewTime(t), nil
}

func (c *RedisClient) HGetFloat64(key, field string) (types.NullFloat64, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Float64(conn.Do("HGET", key, field))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return types.NewFloat64Null(), nil
		} else {
			return types.NewFloat64Null(), err
		}
	}
	return types.NewFloat64(val), nil
}

//HGetAll 对应hgetall命令
func (c *RedisClient) HGetAll(key string) (map[string]interface{}, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	tmp, err := redis.Bytes(conn.Do("HGETALL", key))
	if err != nil {
		log.Error("Command hgetall", zap.Error(err))
		return nil, err
	}

	var val map[string]interface{}

	err = json.Unmarshal(tmp, &val)
	if err != nil {
		log.Error("Json unmarshal failed, ", zap.Error(err))
		return nil, err
	}
	return val, nil
}

//===============================Command Incr===================================

//Incr 对应incr命令
func (c *RedisClient) Incr(key string) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		log.Error("Command incr", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Int(conn.Do("INCRBY", key, n))
	if err != nil {
		log.Error("Command incrby", zap.Error(err))
		return 0, err
	}
	return val, nil
}

//===============================Command Decr===================================

//Decr 对应decr命令
func (c *RedisClient) Decr(key string) (int, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Int(conn.Do("DECR", key))
	if err != nil {
		log.Error("Command desr", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Int(conn.Do("DECRBY", key, n))
	if err != nil {
		log.Error("Command decrby", zap.Error(err))
		return 0, err
	}
	return val, nil
}

// TODO 待测试
//===============================Command SAdd===================================

//SAdd 对应sadd命令
func (c *RedisClient) SAdd(key string, args ...string) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	for _, arg := range args {
		conn.Do("SADD", key, arg)
	}
}

//SMembersStr 对应smembers命令
func (c *RedisClient) SMembersStr(key string) ([]string, error) {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Error("Command smembers str", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Ints(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Error("Command smembers int", zap.Error(err))
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
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Int64s(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Error("Command smembers int64s", zap.Error(err), zap.String("key", key))
		return nil, err
	}
	return val, nil
}

//SISMembers 对应sismembers命令
func (c *RedisClient) SISMembers(key, value string) bool {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Bool(conn.Do("SISMEMBER", key, value))
	if err != nil {
		log.Error(
			"Command sismember",
			zap.Error(err),
			zap.String("key", key),
			zap.String("value", value))
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
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	val, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		log.Error("Command exist", zap.Error(err), zap.String("key", key))
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
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	if _, err := conn.Do("DEL", key); err != nil {
		log.Error("Command del", zap.Error(err), zap.String("key", key))
		return err
	}
	return nil
}

//LPush 对应lpush命令
func (c *RedisClient) LPush(key string, args ...interface{}) error {
	conn := c.pool.Get()

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("Redis conn close failed", zap.Error(err))
		}
	}()

	if _, err := conn.Do("LPUSH", key, args); err != nil {
		log.Error("Command lpush", zap.Error(err), zap.String("key", key), zap.Any("value", args))
		return err
	}
	return nil
}
