package cast

import (
	"strconv"
)

func Any2Bool(v interface{}) (bool, error) {
	if dst, ok := v.(bool); !ok {
		return false, nil
	} else {
		return dst, nil
	}
}

func AnyMust2Bool(v interface{}) bool {
	dst, _ := Any2Bool(v)
	return dst
}

func Str2Bool(v string) (bool, error) {
	return strconv.ParseBool(v)
}
