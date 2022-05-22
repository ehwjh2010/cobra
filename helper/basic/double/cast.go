package double

import (
	"github.com/ehwjh2010/viper/verror"
	"strconv"
)

//=================浮点与interface转化================

func Double2Any(v interface{}) interface{} {
	return v
}

func Any2Double(v interface{}) (float64, error) {
	if dst, ok := v.(float64); !ok {
		return 0, verror.CastDoubleError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Double(v interface{}) float64 {
	dst, _ := Any2Double(v)
	return dst
}

//=================浮点与字符串转化=======================

func Double2Str(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func Str2Double(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func MustStr2Double(v string) float64 {
	dst, _ := Str2Double(v)
	return dst
}

//=================浮点切片与interface转化================

func SliceDouble2Any(vs []float64) []interface{} {

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func SliceAny2Double(vs []interface{}) ([]float64, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]float64, len(vs))
	for i, v := range vs {
		tmp, err := Any2Double(v)
		if err != nil {
			return nil, err
		}

		rs[i] = tmp
	}

	return rs, nil
}
