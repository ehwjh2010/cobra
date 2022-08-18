package integer

import (
	"strconv"

	"github.com/ehwjh2010/viper/verror"
)

// 以下强转未考虑越界情况,
// 可能存在溢出问题的函数分别为 Int64ToInt, IntToInt32, Int64ToInt32

//=======================int与interface转化=====================

func Any2Int(v interface{}) (int, error) {
	if dst, ok := v.(int); !ok {
		return 0, verror.CastIntErr
	} else {
		return dst, nil
	}
}

func MustAny2Int(v interface{}) int {
	dst := v.(int) //nolint:errcheck
	return dst
}

func Int2Any(v interface{}) interface{} {
	return v
}

//=======================int与str转化=====================

func Int2Str(v int) string {
	str := strconv.Itoa(v)
	return str
}

func Int32ToStr(v int32) string {
	return Int64ToStr(int64(v))
}

func Int64ToStr(v int64) string {
	return strconv.FormatInt(v, 10)
}

//=======================int32与interface转化=====================

func Any2Int32(v interface{}) (int32, error) {
	if dst, ok := v.(int32); !ok {
		return 0, verror.CastInt32Err
	} else {
		return dst, nil
	}
}

func MustAny2Int32(v interface{}) int32 {
	dst := v.(int32) //nolint:errcheck
	return dst
}

func Int32TAny(v interface{}) interface{} {
	return v
}

//=======================int64与interface转化=====================

func Any2Int64(v interface{}) (int64, error) {
	if dst, ok := v.(int64); !ok {
		return 0, verror.CastInt64Err
	} else {
		return dst, nil
	}
}

func MustAny2Int64(v interface{}) int64 {
	dst := v.(int64) //nolint:errcheck
	return dst
}

func Int64ToAny(v interface{}) interface{} {
	return v
}

//=======================int与int32转化=========================

func IntToInt32(v int) int32 {
	return int32(v)
}

func Int32ToInt(v int32) int {
	return int(v)
}

//=======================int与int64转化=========================

func IntToInt64(v int) int64 {
	return int64(v)
}

func Int64ToInt(v int64) int {
	return int(v)
}

//=======================int32与int64转化=======================

func Int32ToInt64(v int32) int64 {
	return int64(v)
}

func Int64ToInt32(v int64) int32 {
	return int32(v)
}
