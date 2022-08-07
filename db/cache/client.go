package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ehwjh2010/viper/component/routine"
	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/helper/basic/double"
	"github.com/ehwjh2010/viper/helper/basic/integer"
	"github.com/ehwjh2010/viper/helper/nano"
	"github.com/ehwjh2010/viper/helper/serialize"
	"github.com/ehwjh2010/viper/helper/types"
	"github.com/ehwjh2010/viper/log"
	"github.com/go-redis/redis/v8"
	wrapErr "github.com/pkg/errors"
)

const (
	DelKeyScript = `
        if redis.call("get",KEYS[1]) == ARGV[1] then
            return redis.call("del",KEYS[1])
        else
            return -1
        end`

	RenewExpireScript = `
		local t = redis.call("ttl",KEYS[1])
		local threshold = tonumber(ARGV[1])
		local exp = tonumber(ARGV[2])
		if t == -2 then 
			return -2
		end
		if t <= threshold then
			return redis.call("expire",KEYS[1],exp)
		else
			return 0
		end
		`
)

const (
	KeyNoExists int64 = -2
	NoOperate   int64 = 0
	ReNewExp    int64 = 1
)

var AcquireLockFailed = errors.New("acquire distribute lock failed")

type RedisClient struct {
	// client redis原生client
	client *redis.Client
	// rawConfig 数据库配置配置
	rawConfig *Cache
	// pCount 心跳连续失败次数
	pCount int
	// rCount 重连连续失败次数
	rCount int
}

func NewRedisClient(client *redis.Client, rawConfig *Cache) *RedisClient {
	return &RedisClient{
		client:    client,
		rawConfig: rawConfig,
	}
}

// getExpire 获取过期时间.
func (r *RedisClient) getExpire(ts int) time.Duration {
	var result time.Duration

	if ts <= enums.ZeroSecond {
		result = enums.ZeroSecD
	} else {
		result = time.Duration(ts) * time.Second
	}

	return result
}

// Heartbeat ping连接.
func (r *RedisClient) Heartbeat() error {
	_, err := r.client.Ping(context.TODO()).Result()
	return err
}

// WatchHeartbeat 监测心跳和重连.
func (r *RedisClient) WatchHeartbeat() {
	// TODO 监测逻辑接口化

	fn := func() {
		waitFlag := true
		for {
			if waitFlag {
				<-time.After(3 * time.Second)
			}

			// 重连失败次数大于0, 直接重连
			if r.rCount > 0 {
				if r.rCount >= 3 {
					<-time.After(enums.OneSecD)
				}
				if ok, _ := r.replaceDB(); ok {
					r.rCount = 0
					r.pCount = 0
					waitFlag = true
				} else {
					r.rCount++
					r.pCount++
					waitFlag = false
				}
				continue
			}

			if r.Heartbeat() != nil {
				r.pCount++
				// 心跳连续3次失败, 触发重连
				if r.pCount >= 3 {
					if ok, _ := r.replaceDB(); ok {
						r.rCount = 0
						r.pCount = 0
						waitFlag = true
					} else {
						r.rCount++
						waitFlag = false
					}
				}
			} else {
				r.rCount = 0
				r.pCount = 0
				waitFlag = true
			}
		}
	}

	// 优先使用协程池监听, 如果没有使用原生协程监听
	err := routine.AddTask(fn)
	if err != nil {
		if errors.Is(err, routine.NoEnableRoutinePool) {
			go fn()
		} else {
			log.Warn("watch heartbeat failed")
		}

	}
}

// replaceDB 替换内部client.
func (r *RedisClient) replaceDB() (bool, error) {
	cli, err := initCacheWithGoRedis(r.rawConfig)
	if err != nil {
		return false, err
	}

	// 关闭之前的连接
	r.client.Close()

	r.client = cli
	return true, nil

}

// Close 关闭连接池.
func (r *RedisClient) Close() error {
	err := r.client.Close()
	if err != nil {
		return err
	} else {
		log.Debug("Close redis success!")
		return nil
	}
}

// Exist 确认key是否存在.
func (r *RedisClient) Exist(key string) (bool, error) {
	ctx := context.TODO()

	result, err := r.client.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	exist := false
	if result == 1 {
		exist = true
	}

	return exist, nil
}

// Delete 删除指定key.
func (r *RedisClient) Delete(key ...string) error {
	ctx := context.TODO()

	_, err := r.client.Del(ctx, key...).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	return nil
}

// FlushDB 清空DB.
func (r *RedisClient) FlushDB() error {
	ctx := context.TODO()

	_, err := r.client.FlushDB(ctx).Result()

	return err
}

// AsyncFlushDB 异步清空DB.
func (r *RedisClient) AsyncFlushDB() error {
	ctx := context.TODO()

	_, err := r.client.FlushDBAsync(ctx).Result()

	return err
}

// Expire 设置过期时间.
func (r *RedisClient) Expire(key string, timeoutSeconds int) error {
	ctx := context.TODO()

	_, err := r.client.Expire(ctx, key, time.Duration(timeoutSeconds)*time.Second).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	return nil
}

// ExpireAt 设置过期时间.
func (r *RedisClient) ExpireAt(key string, expireAt time.Time) error {
	ctx := context.TODO()

	_, err := r.client.ExpireAt(ctx, key, expireAt).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	return nil
}

// TTL 获取指定key的过期时间, 返回时间, 单位: s
// 如果结果是nil, 则key不存在
// 如果结果值-1, 则key无过期时间
// 否则值就是key的过期时间.
func (r *RedisClient) TTL(key string) (types.NullFloat64, error) {
	ctx := context.TODO()

	t, err := r.client.TTL(ctx, key).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		return types.NewFloat64Null(), err
	}

	if t == (time.Duration(-1) * time.Nanosecond) {
		return types.NewFloat64(-1), nil
	} else if t == (time.Duration(-2) * time.Nanosecond) {
		return types.NewFloat64Null(), nil
	} else {
		return types.NewFloat64(t.Seconds()), nil
	}
}

// SetNX 存在不操作, 不存在则设置.
func (r *RedisClient) SetNX(key string, value interface{}, timeoutSeconds int) (bool, error) {
	ctx := context.TODO()

	ok, err := r.client.SetNX(ctx, key, value, time.Duration(timeoutSeconds)*time.Second).Result()

	if err != nil {
		return false, err
	}

	return ok, nil
}

//===============================Command Set===================================

// Set redis命令SET, timeoutSeconds <=0, 则认为无过期时间.
func (r *RedisClient) Set(key string, value interface{}, timeoutSeconds int) (err error) {
	ctx := context.TODO()

	expire := r.getExpire(timeoutSeconds)

	if err = r.client.Set(ctx, key, value, expire).Err(); err != nil {
		return wrapErr.Wrap(err, fmt.Sprintf("set key=%s, value=%v err", key, value))
	}

	return nil
}

// SetWithNoExpire redis命令SET, 永久存储.
func (r *RedisClient) SetWithNoExpire(key string, value interface{}) (err error) {
	ctx := context.TODO()

	if err = r.client.Set(ctx, key, value, enums.ZeroSecD).Err(); err != nil {
		return wrapErr.Wrap(err, fmt.Sprintf("set key=%s, value=%v err", key, value))
	}

	return nil
}

// SetJson 设置Json.
func (r *RedisClient) SetJson(key string, value interface{}, timeoutSeconds int) (err error) {

	if value == nil {
		return nil
	}

	str, err := serialize.Marshal(value)
	if err != nil {
		return err
	}

	return r.Set(key, str, timeoutSeconds)
}

// MSet 批量Set.
func (r *RedisClient) MSet(data map[string]interface{}) error {
	ctx := context.TODO()

	return r.client.MSet(ctx, data).Err()
}

//===============================Command Get===================================

// GetStr GetStr.
func (r *RedisClient) GetStr(key string) (types.NullString, error) {
	ctx := context.TODO()

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

// GetInt GetInt.
func (r *RedisClient) GetInt(key string) (types.NullInt, error) {
	ctx := context.TODO()

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

// GetInt64 GetInt64.
func (r *RedisClient) GetInt64(key string) (types.NullInt64, error) {
	ctx := context.TODO()

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

// GetFloat64 GetFloat64.
func (r *RedisClient) GetFloat64(key string) (types.NullFloat64, error) {
	ctx := context.TODO()

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

// GetBool GetBool.
func (r *RedisClient) GetBool(key string) (types.NullBool, error) {
	ctx := context.TODO()

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

// GetTime GetTime.
func (r *RedisClient) GetTime(key string) (types.NullTime, error) {
	ctx := context.TODO()

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

// GetJson GetJson.
func (r *RedisClient) GetJson(key string, dst interface{}) (bool, error) {
	ctx := context.TODO()

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

// Incr 加1.
func (r *RedisClient) Incr(key string) (int64, error) {
	ctx := context.TODO()

	cmd := r.client.Incr(ctx, key)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "incr "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

// IncrBy 添加指定数量.
func (r *RedisClient) IncrBy(key string, incr int64) (int64, error) {
	ctx := context.TODO()

	cmd := r.client.IncrBy(ctx, key, incr)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "incrby "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

// Decr 减1.
func (r *RedisClient) Decr(key string) (int64, error) {
	ctx := context.TODO()

	cmd := r.client.Decr(ctx, key)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "decr "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

// DecrBy 减少指定数量.
func (r *RedisClient) DecrBy(key string, decr int64) (int64, error) {
	ctx := context.TODO()

	cmd := r.client.DecrBy(ctx, key, decr)
	if cmd.Err() != nil {
		return 0, wrapErr.Wrap(cmd.Err(), "decrby "+key+" failed")
	}

	count := cmd.Val()
	return count, nil
}

//===============================Command list===================================

// LPush 往列表插入值.
func (r *RedisClient) LPush(key string, value ...interface{}) error {
	ctx := context.TODO()

	return r.client.LPush(ctx, key, value...).Err()
}

// RPush 往列表插入值.
func (r *RedisClient) RPush(key string, value ...interface{}) error {
	ctx := context.TODO()

	return r.client.RPush(ctx, key, value...).Err()
}

// LMembersStr 获取列表全部内容.
func (r *RedisClient) LMembersStr(key string, start, end int) ([]string, error) {
	ctx := context.TODO()

	result, err := r.client.LRange(ctx, key, int64(start), int64(end)).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberStr 获取列表全部内容.
func (r *RedisClient) LAllMemberStr(key string) ([]string, error) {
	return r.LMembersStr(key, 0, -1)
}

// LMembersInt 获取列表全部内容.
func (r *RedisClient) LMembersInt(key string, start, end int) ([]int, error) {
	ctx := context.TODO()

	result := make([]int, 0)

	err := r.client.LRange(ctx, key, int64(start), int64(end)).ScanSlice(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberInt 获取列表全部内容.
func (r *RedisClient) LAllMemberInt(key string) ([]int, error) {
	return r.LMembersInt(key, 0, -1)
}

// LMembersInt64 获取列表全部内容.
func (r *RedisClient) LMembersInt64(key string, start, end int) ([]int64, error) {
	ctx := context.TODO()

	result := make([]int64, 0)

	err := r.client.LRange(ctx, key, int64(start), int64(end)).ScanSlice(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberInt64 获取列表全部内容.
func (r *RedisClient) LAllMemberInt64(key string) ([]int64, error) {
	return r.LMembersInt64(key, 0, -1)
}

// LMembersFloat64 获取列表全部内容.
func (r *RedisClient) LMembersFloat64(key string, start, end int) ([]float64, error) {
	ctx := context.TODO()

	result := make([]float64, 0)

	err := r.client.LRange(ctx, key, int64(start), int64(end)).ScanSlice(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LAllMemberFloat64 获取列表全部内容.
func (r *RedisClient) LAllMemberFloat64(key string) ([]int64, error) {
	return r.LMembersInt64(key, 0, -1)
}

// LFirstMemberStr 获取第一个元素, 如果值是nil, 代表列表为空或key不存在.
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

// LFirstMemberInt 获取第一个元素, 如果值是nil, 代表列表为空或key不存在.
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

// LPop 从头部删除元素.
func (r *RedisClient) LPop(key string) (types.NullString, error) {
	ctx := context.TODO()

	result, err := r.client.LPop(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewStrNull(), nil
		} else {
			return types.NewStrNull(), err
		}
	}

	return types.NewStr(result), nil
}

// LPopWithCount 从头部删除元素.
func (r *RedisClient) LPopWithCount(key string, count int) ([]string, error) {
	ctx := context.TODO()

	result, err := r.client.LPopCount(ctx, key, count).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	return result, nil
}

// RPop 从尾部删除元素.
func (r *RedisClient) RPop(key string) (types.NullString, error) {
	ctx := context.TODO()

	result, err := r.client.RPop(ctx, key).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		if errors.Is(err, redis.Nil) {
			return types.NewStrNull(), nil
		} else {
			return types.NewStrNull(), err
		}
	}

	return types.NewStr(result), nil
}

// RPopWithCount 从头部删除元素.
func (r *RedisClient) RPopWithCount(key string, count int) ([]string, error) {
	ctx := context.TODO()

	result, err := r.client.RPopCount(ctx, key, count).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	return result, nil
}

// LLen 列表长度.
func (r *RedisClient) LLen(key string) (int64, error) {
	ctx := context.TODO()

	result, err := r.client.LLen(ctx, key).Result()

	return result, err
}

// LRem 删除列表中所有与value相等的元素.
func (r *RedisClient) LRem(key string, value interface{}) error {
	return r.LRemWithCount(key, value, 0)
}

// LRemFirstOne 从头部开始删除第一个value相等的元素.
func (r *RedisClient) LRemFirstOne(key string, value interface{}) error {
	return r.LRemWithCount(key, value, 1)
}

// LRemLastOne 从尾部开始删除第一个value相等的元素.
func (r *RedisClient) LRemLastOne(key string, value interface{}) error {
	return r.LRemWithCount(key, value, -1)
}

// LRemWithCount 删除列表中与value相等的元素, 删除个数为count.
func (r *RedisClient) LRemWithCount(key string, value interface{}, count int) error {
	ctx := context.TODO()

	_, err := r.client.LRem(ctx, key, int64(count), value).Result()

	return err
}

// LTrim 保留指定start, end 范围的元素, 包括边界元素, 其中start, end为列表下标.
func (r *RedisClient) LTrim(key string, start, end int) error {
	ctx := context.TODO()

	_, err := r.client.LTrim(ctx, key, int64(start), int64(end)).Result()

	return err
}

// RPopLPush Redis命令rpoplpush.
func (r *RedisClient) RPopLPush(src, dst string) (string, error) {
	ctx := context.TODO()

	result, err := r.client.RPopLPush(ctx, src, dst).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", wrapErr.Wrap(err, "key="+src+" not exist")
		} else {
			return "", wrapErr.Wrap(err, "operate redis failed")
		}
	}

	return result, nil
}

//===============================Command hash===================================

// HGetStr Redis命令hget.
func (r *RedisClient) HGetStr(key string, field string) (types.NullString, error) {
	ctx := context.TODO()

	value, err := r.client.HGet(ctx, key, field).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewStrNull(), nil
		} else {
			return types.NewStrNull(), err
		}
	}

	return types.NewStr(value), nil
}

// HGetInt Redis命令hget.
func (r *RedisClient) HGetInt(key string, field string) (types.NullInt, error) {
	ctx := context.TODO()

	value, err := r.client.HGet(ctx, key, field).Int()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewIntNull(), nil
		} else {
			return types.NewIntNull(), err
		}
	}

	return types.NewInt(value), nil
}

// HGetInt64 Redis命令hget.
func (r *RedisClient) HGetInt64(key string, field string) (types.NullInt64, error) {
	ctx := context.TODO()

	value, err := r.client.HGet(ctx, key, field).Int64()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewInt64Null(), nil
		} else {
			return types.NewInt64Null(), err
		}
	}

	return types.NewInt64(value), nil
}

// HGetFloat64 Redis命令hget.
func (r *RedisClient) HGetFloat64(key string, field string) (types.NullFloat64, error) {
	ctx := context.TODO()

	value, err := r.client.HGet(ctx, key, field).Float64()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewFloat64Null(), nil
		} else {
			return types.NewFloat64Null(), err
		}
	}

	return types.NewFloat64(value), nil
}

// HGetBool Redis命令hget.
func (r *RedisClient) HGetBool(key string, field string) (types.NullBool, error) {
	ctx := context.TODO()

	value, err := r.client.HGet(ctx, key, field).Bool()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewBoolNull(), nil
		} else {
			return types.NewBoolNull(), err
		}
	}

	return types.NewBool(value), nil
}

// HGetTime Redis命令hget.
func (r *RedisClient) HGetTime(key string, field string) (types.NullTime, error) {
	ctx := context.TODO()

	value, err := r.client.HGet(ctx, key, field).Time()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewTimeNull(), nil
		} else {
			return types.NewTimeNull(), err
		}
	}

	return types.NewTime(value), nil
}

// HGetJson Redis命令hget.
func (r *RedisClient) HGetJson(key string, field string, dst interface{}) (bool, error) {
	ctx := context.TODO()

	v, err := r.client.HGet(ctx, key, field).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		} else {
			return false, err
		}
	}

	if err = serialize.Unmarshal(v, dst); err != nil {
		return false, err
	}

	return true, nil
}

// HGetAll Redis命令hgetall.
func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	ctx := context.TODO()

	v, err := r.client.HGetAll(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return v, nil
}

// HGetAllWIthMap Redis命令hgetall, 返回值第一个是key是否存在, 第二个是错误.
func (r *RedisClient) HGetAllWIthMap(key string) (map[string]string, error) {
	ctx := context.TODO()

	v, err := r.client.HGetAll(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return v, nil
}

// HSet Redis命令Hset.
func (r *RedisClient) HSet(key string, info map[string]interface{}) error {
	ctx := context.TODO()

	if _, err := r.client.HSet(ctx, key, info).Result(); err != nil {
		return err
	}

	return nil
}

// HSetJson Redis命令Hset.
func (r *RedisClient) HSetJson(key, field string, value interface{}) error {
	ctx := context.TODO()

	marshalByte, err := serialize.Marshal(value)
	if err != nil {
		return err
	}

	if _, err := r.client.HSet(ctx, key, map[string]interface{}{field: marshalByte}).Result(); err != nil {
		return err
	}

	return nil
}

// HKeys Redis命令hkeys.
func (r *RedisClient) HKeys(key string) ([]string, error) {
	ctx := context.TODO()

	result, err := r.client.HKeys(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return result, nil
}

// HLen Redis命令hlen.
func (r *RedisClient) HLen(key string) (int64, error) {
	ctx := context.TODO()

	count, err := r.client.HLen(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return count, nil
}

// HDel Redis命令hdel.
func (r *RedisClient) HDel(key string, fields ...string) (int64, error) {
	ctx := context.TODO()

	count, err := r.client.HDel(ctx, key, fields...).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return count, nil
}

// HExists Redis命令hexists.
func (r *RedisClient) HExists(key string, field string) (bool, error) {
	ctx := context.TODO()

	exists, err := r.client.HExists(ctx, key, field).Result()

	if err != nil {
		return false, err
	}

	return exists, nil
}

//===============================Command set===================================

// SAdd Redis命令sadd.
func (r *RedisClient) SAdd(key string, value ...interface{}) error {
	ctx := context.TODO()

	if _, err := r.client.SAdd(ctx, key, value...).Result(); err != nil {
		return err
	}

	return nil
}

// SIsMember Redis命令sismember.
func (r *RedisClient) SIsMember(key string, value interface{}) (bool, error) {
	ctx := context.TODO()

	exist, err := r.client.SIsMember(ctx, key, value).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		} else {
			return false, err
		}
	}

	return exist, nil
}

// SMembers Redis命令smembers.
func (r *RedisClient) SMembers(key string) ([]string, error) {
	ctx := context.TODO()
	result, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return result, nil
}

// SMembersInt Redis命令smembers.
func (r *RedisClient) SMembersInt(key string) ([]int, error) {
	ctx := context.TODO()

	ret := make([]int, 0)

	err := r.client.SMembers(ctx, key).ScanSlice(&ret)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return ret, nil
}

// SMembersInt64 Redis命令smembers.
func (r *RedisClient) SMembersInt64(key string) ([]int64, error) {
	ctx := context.TODO()

	ret := make([]int64, 0)

	err := r.client.SMembers(ctx, key).ScanSlice(&ret)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return ret, nil
}

// SMembersFloat64 Redis命令smembers.
func (r *RedisClient) SMembersFloat64(key string) ([]float64, error) {
	ctx := context.TODO()

	ret := make([]float64, 0)

	err := r.client.SMembers(ctx, key).ScanSlice(&ret)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return ret, nil
}

// SMembersTime Redis命令smembers.
func (r *RedisClient) SMembersTime(key string) ([]time.Time, error) {
	ctx := context.TODO()

	ret := make([]time.Time, 0)

	err := r.client.SMembers(ctx, key).ScanSlice(&ret)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return ret, nil
}

// SPopStr Redis命令spop, 返回删除的值.
func (r *RedisClient) SPopStr(key string) (types.NullString, error) {
	ctx := context.TODO()

	result, err := r.client.SPop(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewStrNull(), nil
		} else {
			return types.NewStrNull(), err
		}
	}

	return types.NewStr(result), nil
}

// SPopInt Redis命令spop.
func (r *RedisClient) SPopInt(key string) (types.NullInt, error) {
	ctx := context.TODO()

	result, err := r.client.SPop(ctx, key).Int()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewIntNull(), nil
		} else {
			return types.NewIntNull(), err
		}
	}

	return types.NewInt(result), nil
}

// SPopInt64 Redis命令spop.
func (r *RedisClient) SPopInt64(key string) (types.NullInt64, error) {
	ctx := context.TODO()

	result, err := r.client.SPop(ctx, key).Int64()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewInt64Null(), nil
		} else {
			return types.NewInt64Null(), err
		}
	}

	return types.NewInt64(result), nil
}

// SPopBool Redis命令spop.
func (r *RedisClient) SPopBool(key string) (types.NullBool, error) {
	ctx := context.TODO()

	result, err := r.client.SPop(ctx, key).Bool()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewBoolNull(), nil
		} else {
			return types.NewBoolNull(), err
		}
	}

	return types.NewBool(result), nil
}

// SPopFloat64 Redis命令spop.
func (r *RedisClient) SPopFloat64(key string) (float64, error) {
	ctx := context.TODO()

	result, err := r.client.SPop(ctx, key).Float64()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return result, nil
}

// SRem Redis命令srem, 返回删除个数.
func (r *RedisClient) SRem(key string, dst ...interface{}) (int64, error) {
	ctx := context.TODO()

	result, err := r.client.SRem(ctx, key, dst...).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return result, nil
}

//===============================Command zset===================================

// ZAdd Redis命令 zadd.
func (r *RedisClient) ZAdd(key string, members ...*redis.Z) error {
	ctx := context.TODO()

	if len(members) <= 0 {
		return nil
	}

	_, err := r.client.ZAdd(ctx, key, members...).Result()
	return err
}

// ZScore Redis命令zscore.
func (r *RedisClient) ZScore(key string, value string) (types.NullFloat64, error) {
	ctx := context.TODO()

	result, err := r.client.ZScore(ctx, key, value).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return types.NewFloat64Null(), nil
		} else {
			return types.NewFloat64Null(), err
		}
	}

	return types.NewFloat64(result), nil
}

// ZCount Redis命令zcount, 返回score 值在 min 和 max 之间的成员的数量.
func (r *RedisClient) ZCount(key string, scoreMin, scoreMax string) (int64, error) {
	ctx := context.TODO()

	result, err := r.client.ZCount(ctx, key, scoreMin, scoreMax).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return result, nil
}

func (r *RedisClient) ZCountWithFloat64(key string, scoreMin, scoreMax float64) (int64, error) {

	min := double.Double2Str(scoreMin)
	max := double.Double2Str(scoreMax)
	return r.ZCount(key, min, max)
}

func (r *RedisClient) ZCountWithInt(key string, scoreMin, scoreMax int) (int64, error) {
	min := integer.Int2Str(scoreMin)
	max := integer.Int2Str(scoreMax)
	return r.ZCount(key, min, max)
}

func (r *RedisClient) ZCountWithInt32(key string, scoreMin, scoreMax int32) (int64, error) {
	min := integer.Int32ToStr(scoreMin)
	max := integer.Int32ToStr(scoreMax)
	return r.ZCount(key, min, max)
}

func (r *RedisClient) ZCountWithInt64(key string, scoreMin, scoreMax int64) (int64, error) {
	min := integer.Int64ToStr(scoreMin)
	max := integer.Int64ToStr(scoreMax)
	return r.ZCount(key, min, max)
}

// ZRange Redis命令zrange, 包括start, end 边界值,  返回按照score升序.
func (r *RedisClient) ZRange(key string, start, end int64) ([]string, error) {
	ctx := context.TODO()

	result, err := r.client.ZRange(ctx, key, start, end).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return result, nil
}

// ZRangeWithCount 返回按照score升序.
func (r *RedisClient) ZRangeWithCount(key string, count int64) ([]string, error) {
	return r.ZRange(key, 0, count-1)
}

// ZRevRange Redis命令zrevrange, 包括start, end 边界值,  返回按照score降序.
func (r *RedisClient) ZRevRange(key string, start, end int64) ([]string, error) {
	ctx := context.TODO()

	strings, err := r.client.ZRevRange(ctx, key, start, end).Result()

	if err != nil {
		return nil, err
	}

	return strings, nil

}

// ZRevRangeWithCount Redis命令zrevrange, 返回按照score降序.
func (r *RedisClient) ZRevRangeWithCount(key string, count int64) ([]string, error) {
	return r.ZRevRange(key, 0, count-1)
}

// ZRank Redis命令zrank, 返回按照score升序, 返回排名.
func (r *RedisClient) ZRank(key string, member string) (int64, error) {
	ctx := context.TODO()

	score, err := r.client.ZRank(ctx, key, member).Result()

	if err != nil {
		return 0, err
	}

	return score, nil
}

// ZRevRank Redis命令zrevrank, 按照score降序, 返回排名.
func (r *RedisClient) ZRevRank(key string, member string) (int64, error) {
	ctx := context.TODO()

	score, err := r.client.ZRevRank(ctx, key, member).Result()

	if err != nil {
		return 0, err
	}

	return score, nil
}

// ZRangeByScore Redis命令zrange, 包括start, end 边界值, 返回按照score排序.
func (r *RedisClient) ZRangeByScore(key string, scoreMin, scoreMax string, reverse bool) ([]redis.Z, error) {
	ctx := context.TODO()

	z := redis.ZRangeArgs{
		Key:     key,
		Start:   scoreMin,
		Stop:    scoreMax,
		ByScore: true,
		Rev:     reverse,
	}

	result, err := r.client.ZRangeArgsWithScores(ctx, z).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return result, nil
}

// ZRemRangeByRank Redis命令ZREMRANGEBYRANK 根据排名删除 按照score升序.
func (r *RedisClient) ZRemRangeByRank(key string, start, end int64) (uint64, error) {
	ctx := context.TODO()

	count, err := r.client.ZRemRangeByRank(ctx, key, start, end).Uint64()

	if err != nil {
		return 0, err
	}

	return count, nil
}

// ZRemRangeByRankCount 根据排名升序删除 排名按照score升序.
func (r *RedisClient) ZRemRangeByRankCount(key string, count int64) (uint64, error) {
	return r.ZRemRangeByRank(key, 0, count-1)
}

// ZRemRangeByScore Redis命令ZREMRANGEBYSCORE.
func (r *RedisClient) ZRemRangeByScore(key string, min, max string) (uint64, error) {
	ctx := context.TODO()

	count, err := r.client.ZRemRangeByScore(ctx, key, min, max).Uint64()

	if err != nil {
		return 0, err
	}

	return count, nil
}

// ZRem Redis命令zrem, 删除指定member的field.
func (r *RedisClient) ZRem(key string, members ...interface{}) (int64, error) {
	ctx := context.TODO()

	result, err := r.client.ZRem(ctx, key, members...).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return result, nil
}

// Pipeline 获取管道.
func (r *RedisClient) Pipeline() redis.Pipeliner {
	pipeliner := r.client.Pipeline()
	return pipeliner
}

// Scan 扫描.
func (r *RedisClient) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	ctx := context.TODO()
	keys, cur, err := r.client.Scan(ctx, cursor, match, count).Result()
	return keys, cur, err
}

// ScanPattern 扫描key, 扫描整张表.
func (r *RedisClient) ScanPattern(match string, count int64, ch chan<- string, sleepSeconds int) {
	var cursor uint64
	sleepFlag := sleepSeconds > 0

	for {
		strings, temp, err := r.Scan(cursor, match, count)
		if err != nil {
			close(ch)
			panic(err)
		}

		cursor = temp
		for _, str := range strings {
			ch <- str
		}

		if cursor == 0 {
			close(ch)
			break
		}

		if sleepFlag {
			time.Sleep(time.Duration(sleepSeconds) * time.Second)
		}
	}
}

// SetArgs 设置
// 如果设置了NX, 并且key已存在, 返回redis.Nil错误
// 成功设置返回OK.
func (r *RedisClient) SetArgs(key string, value interface{}, args redis.SetArgs) (string, error) {
	ctx := context.TODO()
	result, err := r.client.SetArgs(ctx, key, value, args).Result()
	return result, err
}

// GetClient 获取原生client.
func (r *RedisClient) GetClient() *redis.Client {
	cli := r.client
	return cli
}

// DoWithDistributeLock 使用分布式锁执行任务.
func (r *RedisClient) DoWithDistributeLock(req *DistributeKeyParam) error {
	expireAt := time.Now().Add(req.AcquireTime).Unix()
	value := nano.MustGetNanoId()

	var acquire bool
	for time.Now().Unix() <= expireAt {
		if _, err := r.SetArgs(req.Key, value, redis.SetArgs{
			Mode: "NX",
			TTL:  req.LockTime,
		}); err != nil {
			if errors.Is(err, redis.Nil) {
				time.Sleep(time.Duration(100) * time.Millisecond)
				continue
			}

			return err
		}

		acquire = true
		break
	}

	if !acquire {
		return AcquireLockFailed
	}

	ch := make(chan struct{})
	go r.scheduleDistributeLock(req.Key, req.Period, req.LockTime, ch)
	defer func() {
		ch <- struct{}{}
		ctx := context.TODO()
		r.client.Eval(ctx, DelKeyScript, []string{req.Key}, value)
	}()

	return req.Fn()
}

// scheduleDistributeLock 添加定时任务, 定时给分布式Key重置过期时间.
func (r *RedisClient) scheduleDistributeLock(key string, period time.Duration, expire time.Duration, stopChan <-chan struct{}) {
	t := time.NewTicker(period)

	seconds := int64(expire.Seconds())
	renewThreshold := seconds / 2
	values := []interface{}{+renewThreshold, +seconds}

	for {
		select {
		case <-t.C:
			ctx := context.TODO()
			result, err := r.client.Eval(ctx, RenewExpireScript, []string{key}, values...).Result()
			if err != nil {
				log.Errorf("renew expire failed, err=%v", err)
			} else {
				ret := integer.MustAny2Int64(result)
				switch ret {
				case KeyNoExists:
					log.Debug("distribute key no exists")
					return
				case NoOperate:
					log.Debug("no operate distribute key")
				case ReNewExp:
					log.Debug("renew expire")
				}
			}
		case <-stopChan:
			log.Debug("receive stop signal")
			return
		}
	}
}
