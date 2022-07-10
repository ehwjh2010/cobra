package str

import (
	"github.com/ehwjh2010/viper/verror"
	"go.uber.org/zap/buffer"
	"strconv"
)

//=================str与interface转化==================

func Char2Any(v string) interface{} {
	return v
}

func Any2Char(v interface{}) (string, error) {
	if dst, ok := v.(string); !ok {
		return "", verror.CastStringError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Char(v interface{}) string {
	dst, _ := Any2Char(v)
	return dst
}

func SliceStr2Any(vs []string) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func SliceAny2Char(vs []interface{}) ([]string, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]string, len(vs))
	for i, v := range vs {
		value, err := Any2Char(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}

//=================str与bytes转化==================

func Char2Bytes(v string) []byte {
	var buff buffer.Buffer

	_, _ = buff.WriteString(v)

	return buff.Bytes()
}

func Char2Int(v string) (int, error) {
	return strconv.Atoi(v)
}

func MustChar2Int(v string) int {
	dst, _ := Char2Int(v)
	return dst
}

func Char2Int32(v string) (int32, error) {
	parseInt, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(parseInt), err
}

func MustChar2Int32(v string) int32 {
	dst, _ := Char2Int32(v)
	return dst
}

func Char2Int64(v string) (int64, error) {
	parseInt, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}

	return parseInt, err
}

func MustChar2Int64(v string) int64 {
	dst, _ := Char2Int64(v)
	return dst
}
