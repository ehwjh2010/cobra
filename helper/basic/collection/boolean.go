package collection

import "github.com/ehwjh2010/viper/helper/basic/boolean"

func BoolSlice2AnySlice(v []bool) []interface{} {
	dst := make([]interface{}, len(v))
	for idx, val := range v {
		dst[idx] = val
	}
	return dst
}

func AnySlice2BoolSlice(v []interface{}) ([]bool, error) {
	dst := make([]bool, len(v))
	for idx, val := range v {
		temp, err := boolean.Any2Bool(val)
		if err != nil {
			return nil, err
		}

		dst[idx] = temp
	}
	return dst, nil
}
