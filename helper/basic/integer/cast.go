package integer

import (
	"github.com/ehwjh2010/viper/helper/basic/collection"
	"github.com/ehwjh2010/viper/verror"
	"strconv"
)

// 以下强转未考虑越界情况,
// 可能存在溢出问题的函数分别为 Int64ToInt, IntToInt32, Int64ToInt32

//=======================int与interface转化=====================

func Any2Int(v interface{}) (int, error) {
	if dst, ok := v.(int); !ok {
		return 0, verror.CastIntegerError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Int(v interface{}) int {
	dst := v.(int)
	return dst
}

func Int2Any(v interface{}) interface{} {
	return v
}

func IntSlice2Any(vs []int) []interface{} {
	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
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

func MustAnySlice2Int(vs []interface{}) []int {
	if vs == nil {
		return nil
	}

	rs := make([]int, len(vs))
	for i, v := range vs {
		value, _ := Any2Int(v)
		rs[i] = value
	}

	return rs
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
		return 0, verror.CastIntegerError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Int32(v interface{}) int32 {
	dst := v.(int32)
	return dst
}

func Int32TAny(v interface{}) interface{} {
	return v
}

func SliceInt32TAny(vs []int32) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func SliceAny2Int32(vs []interface{}) ([]int32, error) {
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

func MustAnySlice2Int32(vs []interface{}) []int32 {
	if vs == nil {
		return nil
	}

	rs := make([]int32, len(vs))
	for i, v := range vs {
		value := v.(int32)
		rs[i] = value
	}

	return rs
}

//=======================int64与interface转化=====================

func Any2Int64(v interface{}) (int64, error) {
	if dst, ok := v.(int64); !ok {
		return 0, verror.CastIntegerError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Int64(v interface{}) int64 {
	dst := v.(int64)
	return dst
}

func Int64TAny(v interface{}) interface{} {
	return v
}

func SliceInt64TAny(vs []int64) []interface{} {
	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = Int64TAny(v)
	}

	return rs
}

func SliceAny2Int64(vs []interface{}) ([]int64, error) {
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

func MustAnySlice2Int64(vs []interface{}) []int64 {
	if vs == nil {
		return nil
	}

	rs := make([]int64, len(vs))
	for i, v := range vs {
		value := v.(int64)
		rs[i] = value
	}

	return rs
}

//=======================int与int32转化=========================

func IntToInt32(v int) int32 {
	return int32(v)
}

func SliceInt2Int32(vs []int) []int32 {
	if collection.IsEmptyInt(vs) {
		return nil
	}

	result := make([]int32, len(vs))
	for idx, v := range vs {
		result[idx] = int32(v)
	}

	return result
}

func Int32ToInt(v int32) int {
	return int(v)
}

func SliceInt32Int(vs []int32) []int {
	if collection.IsEmptyInt32(vs) {
		return nil
	}

	result := make([]int, len(vs))
	for idx, v := range vs {
		result[idx] = int(v)
	}

	return result
}

//=======================int与int64转化=========================

func IntToInt64(v int) int64 {
	return int64(v)
}

func SliceInt2Int64(vs []int) []int64 {
	if collection.IsEmptyInt(vs) {
		return nil
	}

	result := make([]int64, len(vs))
	for idx, v := range vs {
		result[idx] = int64(v)
	}

	return result
}

func Int64ToInt(v int64) int {
	return int(v)
}

func SliceInt64TInt(vs []int64) []int {
	if collection.IsEmptyInt64(vs) {
		return nil
	}

	result := make([]int, len(vs))
	for idx, v := range vs {
		result[idx] = int(v)
	}

	return result
}

//=======================int32与int64转化=======================

func Int32ToInt64(v int32) int64 {
	return int64(v)
}

func SliceInt32TInt64(vs []int32) []int64 {
	if collection.IsEmptyInt32(vs) {
		return nil
	}

	result := make([]int64, len(vs))
	for idx, v := range vs {
		result[idx] = int64(v)
	}

	return result
}

func Int64ToInt32(v int64) int32 {
	return int32(v)
}

func SliceInt64TInt32(vs []int64) []int32 {
	if collection.IsEmptyInt64(vs) {
		return nil
	}

	result := make([]int32, len(vs))
	for idx, v := range vs {
		result[idx] = int32(v)
	}

	return result
}
