package collection

func IsEmptyAny(v []interface{}) bool {
	return len(v) <= 0
}

func IsEmptyStr(v []string) bool {
	return len(v) <= 0
}

func IsEmptyBytes(v []byte) bool {
	return len(v) <= 0
}

func IsEmptyInt(v []int) bool {
	return len(v) <= 0
}

func IsEmptyInt32(v []int32) bool {
	return len(v) <= 0
}

func IsEmptyInt64(v []int64) bool {
	return len(v) <= 0
}

func IsEmptyDouble(v []float64) bool {
	return len(v) <= 0
}
