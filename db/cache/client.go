package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/ehwjh2010/cobra/types"
	"github.com/ehwjh2010/cobra/util/serialize"
	"github.com/ehwjh2010/cobra/util/timer"
	"github.com/go-redis/redis/v8"
	wrapErr "github.com/pkg/errors"
	"time"
)

type RedisClient struct {
	client         *redis.Client
	defaultTimeOut int
}

func NewRedisClient(client *redis.Client, defaultTimeOut int) *RedisClient {
	return &RedisClient{client: client, defaultTimeOut: defaultTimeOut}
}

// Close 关闭连接池
func (r *RedisClient) Close() error {
	return r.client.Close()
}

//===============================Command Set===================================

// Set redis命令SET, exp 单位: 秒
func (r *RedisClient) Set(key string, value interface{}, exp int) (err error) {
	ctx := context.Background()

	if exp <= 0 {
		exp = r.defaultTimeOut
	}

	if err = r.client.Set(ctx, key, value, time.Duration(exp)*time.Second).Err(); err != nil {
		return wrapErr.Wrap(err, fmt.Sprintf("set key=%s, value=%v err", key, value))
	}

	return nil
}

//SetNoExpire redis命令SET, 没有过期时间
func (r *RedisClient) SetNoExpire(key string, value interface{}) (err error) {
	ctx := context.Background()

	if err = r.client.Set(ctx, key, value, 0).Err(); err != nil {
		return wrapErr.Wrap(err, fmt.Sprintf("set key=%s, value=%v err", key, value))
	}

	return nil
}

//SetTime 设置时间
func (r *RedisClient) SetTime(key string, value time.Time, exp int) (err error) {
	str := timer.Time2Str(value)

	return r.Set(key, str, exp)
}

//SetJson 设置Json
func (r *RedisClient) SetJson(key string, value interface{}, exp int) (err error) {

	str, err := serialize.MarshalStr(value)
	if err != nil {
		return err
	}

	return r.Set(key, str, exp)
}

//MSet 批量Set
func (r *RedisClient) MSet(data map[string]interface{}) error {
	ctx := context.Background()

	return r.client.MSet(ctx, data).Err()
}

//===============================Command Get===================================

//GetStr GetStr
func (r *RedisClient) GetStr(key string) (types.NullString, error) {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewStrNull(), nil
		} else {
			return types.NewStrNull(), err
		}
	}

	return types.NewStr(result), nil
}

//GetInt GetInt
func (r *RedisClient) GetInt(key string) (types.NullInt, error) {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Int()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewIntNull(), nil
		} else {
			return types.NewIntNull(), err
		}
	}

	return types.NewInt(result), nil
}

//GetInt64 GetInt64
func (r *RedisClient) GetInt64(key string) (types.NullInt64, error) {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Int64()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewInt64Null(), nil
		} else {
			return types.NewInt64Null(), err
		}
	}

	return types.NewInt64(result), nil
}

//GetFloat64 GetFloat64
func (r *RedisClient) GetFloat64(key string) (types.NullFloat64, error) {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Float64()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewFloat64Null(), nil
		} else {
			return types.NewFloat64Null(), err
		}
	}

	return types.NewFloat64(result), nil
}

//GetBool GetBool
func (r *RedisClient) GetBool(key string) (types.NullBool, error) {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Bool()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewBoolNull(), nil
		} else {
			return types.NewBoolNull(), err
		}
	}

	return types.NewBool(result), nil
}

//GetTime GetTime TODO 待测试
func (r *RedisClient) GetTime(key string) (types.NullTime, error) {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Time()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewTimeNull(), nil
		} else {
			return types.NewTimeNull(), err
		}
	}

	return types.NewTime(result), nil
}

//GetJson GetJson TODO 待测试
func (r *RedisClient) GetJson(key string, dst interface{}) error {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Bytes()

	if err != nil {
		return err
	}

	return serialize.Unmarshal(result, dst)
}

//===============================Command Count===================================

//Incr Incr
func (r *RedisClient) Incr(key string) (int64, error) {
	ctx := context.Background()

	cmd := r.client.Incr(ctx, key)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "incr "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

//IncrBy IncrBy
func (r *RedisClient) IncrBy(key string, incr int64) (int64, error) {
	ctx := context.Background()

	cmd := r.client.IncrBy(ctx, key, incr)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "incrby "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

//Decr Decr
func (r *RedisClient) Decr(key string) (int64, error) {
	ctx := context.Background()

	cmd := r.client.Decr(ctx, key)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "decr "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

//DecrBy DecrBy
func (r *RedisClient) DecrBy(key string, decr int64) (int64, error) {
	ctx := context.Background()

	cmd := r.client.DecrBy(ctx, key, decr)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "decrby "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

//===============================Command list===================================

//LPush 往列表插入值
func (r *RedisClient) LPush(key string, value ...interface{}) error {
	ctx := context.Background()

	return r.client.LPush(ctx, key, value...).Err()
}

//RPush 往列表插入值
func (r *RedisClient) RPush(key string, value ...interface{}) error {
	ctx := context.Background()

	return r.client.RPush(ctx, key, value...).Err()
}

//LAllMemberStr 获取列表全部内容
func (r *RedisClient) LAllMemberStr(key string) ([]string, error) {
	ctx := context.Background()

	result, err := r.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

//LPop 从末端删除元素
func (r *RedisClient) LPop(key string) (string, error) {
	ctx := context.Background()

	result, err := r.client.LPop(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return result, nil
}
