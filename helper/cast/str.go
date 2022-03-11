package cast

import (
	"github.com/ehwjh2010/viper/client/verror"
	"go.uber.org/zap/buffer"
)

func Any2String(v interface{}) (string, error) {
	if dst, ok := v.(string); !ok {
		return "", verror.CastStringError(v)
	} else {
		return dst, nil
	}
}

func MustAny2String(v interface{}) string {
	dst, _ := Any2String(v)
	return dst
}

func Str2Bytes(v string) []byte {
	var buff buffer.Buffer

	_, _ = buff.WriteString(v)

	return buff.Bytes()
}

func Str2Any(v string) interface{} {
	return v
}

func StrSlice2Any(vs []string) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func AnySlice2Str(vs []interface{}) ([]string, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]string, len(vs))
	for i, v := range vs {
		value, err := Any2String(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}
