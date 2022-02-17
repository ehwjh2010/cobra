package cast

import (
	"github.com/ehwjh2010/viper/client/verror"
	"strconv"
)

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

func Str2Double(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func MustStr2Double(v string) float64 {
	dst, _ := Str2Double(v)
	return dst
}

func Double2Str(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}
