package boolean

import (
	"github.com/ehwjh2010/viper/verror"
	"strconv"
)

//=================bool与interface转化================

func Any2Bool(v interface{}) (bool, error) {
	if dst, ok := v.(bool); !ok {
		return false, verror.CastBoolError(v)
	} else {
		return dst, nil
	}
}

func MustAny2Bool(v interface{}) bool {
	dst, _ := v.(bool)
	return dst
}

func SliceBool2Any(v []bool) []interface{} {
	dst := make([]interface{}, len(v))
	for idx, val := range v {
		dst[idx] = val
	}
	return dst
}

//=================str与interface转化==================

func Str2Bool(v string) (bool, error) {
	return strconv.ParseBool(v)
}

func MustStr2Bool(v string) bool {
	dst, _ := Str2Bool(v)
	return dst
}

func Bool2Str(b bool) string {
	return strconv.FormatBool(b)
}
