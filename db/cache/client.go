package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/types"
	"github.com/ehwjh2010/cobra/util/serialize"
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
	err := r.client.Close()
	if err != nil {
		return err
	} else {
		log.Debug("Close redis success!")
		return nil
	}
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

//SetWithNoExpire redis命令SET, 永久存储
func (r *RedisClient) SetWithNoExpire(key string, value interface{}) (err error) {
	ctx := context.Background()

	if err = r.client.Set(ctx, key, value, 0).Err(); err != nil {
		return wrapErr.Wrap(err, fmt.Sprintf("set key=%s, value=%v err", key, value))
	}

	return nil
}

//SetJson 设置Json, exp 单位: 秒
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

//GetTime GetTime
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

//GetJson GetJson
func (r *RedisClient) GetJson(key string, dst interface{}) (bool, error) {
	ctx := context.Background()

	result, err := r.client.Get(ctx, key).Bytes()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		} else {
			return false, err
		}
	}

	if err = serialize.Unmarshal(result, dst); err != nil {
		return false, err
	}

	return true, nil
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

// LMembersStr 获取列表全部内容
func (r *RedisClient) LMembersStr(key string, start, end int) ([]string, error) {
	ctx := context.Background()

	result, err := r.client.LRange(ctx, key, int64(start), int64(end)).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberStr 获取列表全部内容
func (r *RedisClient) LAllMemberStr(key string) ([]string, error) {
	return r.LMembersStr(key, 0, -1)
}

// LMembersInt 获取列表全部内容
func (r *RedisClient) LMembersInt(key string, start, end int) ([]int, error) {
	ctx := context.Background()

	result := make([]int, 0)

	err := r.client.LRange(ctx, key, int64(start), int64(end)).ScanSlice(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberInt 获取列表全部内容
func (r *RedisClient) LAllMemberInt(key string) ([]int, error) {
	return r.LMembersInt(key, 0, -1)
}

// LMembersInt64 获取列表全部内容
func (r *RedisClient) LMembersInt64(key string, start, end int) ([]int64, error) {
	ctx := context.Background()

	result := make([]int64, 0)

	err := r.client.LRange(ctx, key, int64(start), int64(end)).ScanSlice(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberInt64 获取列表全部内容
func (r *RedisClient) LAllMemberInt64(key string) ([]int64, error) {
	return r.LMembersInt64(key, 0, -1)
}

// LMembersFloat64 获取列表全部内容
func (r *RedisClient) LMembersFloat64(key string, start, end int) ([]float64, error) {
	ctx := context.Background()

	result := make([]float64, 0)

	err := r.client.LRange(ctx, key, int64(start), int64(end)).ScanSlice(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberFloat64 获取列表全部内容
func (r *RedisClient) LAllMemberFloat64(key string) ([]int64, error) {
	return r.LMembersInt64(key, 0, -1)
}

// TODO 待测试, 核心测试: key不存在时, 会是什么结果

func (r *RedisClient) LFirstMemberStr(key string) (types.NullString, error) {

	result, err := r.LMembersStr(key, 0, 0)

	if err != nil {
		return types.NewStrNull(), err
	}

	if len(result) > 0 {
		return types.NewStr(result[0]), nil
	}

	return types.NewStrNull(), nil
}

func (r *RedisClient) LFirstMemberInt(key string) (types.NullInt, error) {

	result, err := r.LMembersInt(key, 0, 0)

	if err != nil {
		return types.NewIntNull(), err
	}

	if len(result) > 0 {
		return types.NewInt(result[0]), nil
	}

	return types.NewIntNull(), nil
}

func (r *RedisClient) LFirstMemberInt64(key string) (types.NullInt64, error) {

	result, err := r.LMembersInt64(key, 0, 0)

	if err != nil {
		return types.NewInt64Null(), err
	}

	if len(result) > 0 {
		return types.NewInt64(result[0]), nil
	}

	return types.NewInt64Null(), nil
}

func (r *RedisClient) LFirstMemberFloat64(key string) (types.NullFloat64, error) {

	result, err := r.LMembersFloat64(key, 0, 0)

	if err != nil {
		return types.NewFloat64Null(), err
	}

	if len(result) > 0 {
		return types.NewFloat64(result[0]), nil
	}

	return types.NewFloat64Null(), nil
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

//LLen 列表长度
func (r *RedisClient) LLen(key string) (int64, error) {
	ctx := context.Background()

	result, err := r.client.LLen(ctx, key).Result()

	return result, err
}

//LRem 删除列表中所有与value相等的元素
func (r *RedisClient) LRem(key string, value interface{}) error {
	return r.LRemWithCount(key, value, 0)
}

//LRemFirstOne 从头部开始删除第一个value相等的元素
func (r *RedisClient) LRemFirstOne(key string, value interface{}) error {
	return r.LRemWithCount(key, value, 1)
}

//LRemLastOne 从尾部开始删除第一个value相等的元素
func (r *RedisClient) LRemLastOne(key string, value interface{}) error {
	return r.LRemWithCount(key, value, -1)
}

//LRemWithCount 删除列表中与value相等的元素, 删除个数为count
func (r *RedisClient) LRemWithCount(key string, value interface{}, count int) error {
	ctx := context.Background()

	_, err := r.client.LRem(ctx, key, int64(count), value).Result()

	return err
}

// LTrim 保留指定start, end 范围的元素
func (r *RedisClient) LTrim(key string, start, end int) error {
	ctx := context.Background()

	_, err := r.client.LTrim(ctx, key, int64(start), int64(end)).Result()

	return err
}

//RPopLPush Redis命令rpoplpush
func (r *RedisClient) RPopLPush(src ,dst string) (string, error) {
	ctx := context.Background()

	return r.client.RPopLPush(ctx, src, dst).Result()
}
