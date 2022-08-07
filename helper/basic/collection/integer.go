package collection

import (
	"github.com/ehwjh2010/viper/helper/basic/integer"
)

func IntSlice2AnySlice(vs []int) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func AnySlice2IntSlice(vs []interface{}) ([]int, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]int, len(vs))
	for i, v := range vs {
		value, err := integer.Any2Int(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}

func MustAnySlice2IntSlice(vs []interface{}) []int {
	if vs == nil {
		return nil
	}

	rs := make([]int, len(vs))
	for i, v := range vs {
		value, _ := integer.Any2Int(v)
		rs[i] = value
	}

	return rs
}

func Int32Slice2AnySlice(vs []int32) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func AnySlice2Int32Slice(vs []interface{}) ([]int32, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]int32, len(vs))
	for i, v := range vs {
		value, err := integer.Any2Int32(v)
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
		value := v.(int32) //nolint:errcheck
		rs[i] = value
	}

	return rs
}

func Int64Slice2AnySlice(vs []int64) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = integer.Int64ToAny(v)
	}

	return rs
}

func AnySlice2Int64Slice(vs []interface{}) ([]int64, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]int64, len(vs))
	for i, v := range vs {
		value, err := integer.Any2Int64(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}

func MustAnySlice2Int64Slice(vs []interface{}) []int64 {
	if vs == nil {
		return nil
	}

	rs := make([]int64, len(vs))
	for i, v := range vs {
		value := v.(int64) //nolint:errcheck
		rs[i] = value
	}

	return rs
}

func IntSlice2Int32Slice(vs []int) []int32 {
	if vs == nil {
		return nil
	}

	result := make([]int32, len(vs))
	for idx, v := range vs {
		result[idx] = int32(v)
	}

	return result
}

func Int32Slice2IntSlice(vs []int32) []int {
	if vs == nil {
		return nil
	}

	result := make([]int, len(vs))
	for idx, v := range vs {
		result[idx] = int(v)
	}

	return result
}

func IntSlice2Int64Slice(vs []int) []int64 {
	if vs == nil {
		return nil
	}

	result := make([]int64, len(vs))
	for idx, v := range vs {
		result[idx] = int64(v)
	}

	return result
}

func Int64Slice2IntSlice(vs []int64) []int {
	if vs == nil {
		return nil
	}

	result := make([]int, len(vs))
	for idx, v := range vs {
		result[idx] = int(v)
	}

	return result
}

func Int32Slice2Int64Slice(vs []int32) []int64 {
	if vs == nil {
		return nil
	}

	result := make([]int64, len(vs))
	for idx, v := range vs {
		result[idx] = int64(v)
	}

	return result
}

func Int64Slice2Int32Slice(vs []int64) []int32 {
	if vs == nil {
		return nil
	}

	result := make([]int32, len(vs))
	for idx, v := range vs {
		result[idx] = int32(v)
	}

	return result
}
