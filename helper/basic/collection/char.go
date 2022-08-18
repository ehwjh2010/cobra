package collection

import "github.com/ehwjh2010/viper/helper/basic/str"

func CharSlice2AnySlice(vs []string) []interface{} {
	if vs == nil {
		return nil
	}

	rs := make([]interface{}, len(vs))
	for i, v := range vs {
		rs[i] = v
	}

	return rs
}

func AnySlice2CharSlice(vs []interface{}) ([]string, error) {
	if vs == nil {
		return nil, nil
	}

	rs := make([]string, len(vs))
	for i, v := range vs {
		value, err := str.Any2Char(v)
		if err != nil {
			return nil, err
		}
		rs[i] = value
	}

	return rs, nil
}

func MustAnySlice2CharSlice(vs []interface{}) []string {
	if vs == nil {
		return nil
	}

	rs := make([]string, len(vs))
	for i, v := range vs {
		value := str.MustAny2Char(v)
		rs[i] = value
	}

	return rs
}
