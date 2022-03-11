package cast

import (
	"github.com/ehwjh2010/viper/client/verror"
	"strconv"
)

// 以下强转未考虑越界情况,
// 可能存在溢出问题的函数分别为 Int64ToInt, IntToInt32, Int64ToInt32

func Int32ToInt(v int32) int {
	return int(v)
}

func Int64ToInt(v int64) int {
	return int(v)
}

func IntToInt32(v int) int32 {
	return int32(v)
}

func Int64ToInt32(v int64) int32 {
	return int32(v)
}

func IntToInt64(v int) int64 {
	return int64(v)
}

func Int32ToInt64(v int32) int64 {
	return int64(v)
}

func Any2Int32(v interface{}) (int32, error) {
	if dst, ok := v.(int32); !ok {
		return 0, verror.CastIntegerError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Int32(v interface{}) int32 {
	dst, _ := Any2Int32(v)
	return dst
}

func Any2Int64(v interface{}) (int64, error) {
	if dst, ok := v.(int64); !ok {
		return 0, verror.CastIntegerError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Int64(v interface{}) int64 {
	dst, _ := Any2Int64(v)
	return dst
}

func Any2Int(v interface{}) (int, error) {
	if dst, ok := v.(int); !ok {
		return 0, verror.CastIntegerError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Int(v interface{}) int {
	dst, _ := Any2Int(v)
	return dst
}

func Str2Int(v string) (int, error) {
	return strconv.Atoi(v)
}

func MustStr2Int(v string) int {
	dst, _ := Str2Int(v)
	return dst
}

func Str2Int32(v string) (int32, error) {
	parseInt, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(parseInt), err
}

func MustStr2Int32(v string) int32 {
	dst, _ := Str2Int32(v)
	return dst
}

func Str2Int64(v string) (int64, error) {
	parseInt, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}

	return parseInt, err
}

func MustStr2Int64(v string) int64 {
	dst, _ := Str2Int64(v)
	return dst
}

func Integer2Any(v interface{}) interface{} {
	return v
}

func IntSlice2Any(vs []int) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func Int32Slice2Any(vs []int32) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func Int64Slice2Any(vs []int64) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func AnySlice2Int64(vs []interface{}) ([]int64, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]int64, len(vs))
	for i, v := range vs {
		value, err := Any2Int64(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}

func AnySlice2Int32(vs []interface{}) ([]int32, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]int32, len(vs))
	for i, v := range vs {
		value, err := Any2Int32(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}

func AnySlice2Int(vs []interface{}) ([]int, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]int, len(vs))
	for i, v := range vs {
		value, err := Any2Int(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}
