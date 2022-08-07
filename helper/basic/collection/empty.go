package collection

func IsEmptyAny(v []interface{}) bool {
	return len(v) <= 0
}

func IsNotEmptyAny(v []interface{}) bool {
	return !IsEmptyAny(v)
}

func IsEmptyStr(v []string) bool {
	return len(v) <= 0
}

func IsNotEmptyStr(v []string) bool {
	return !IsEmptyStr(v)
}

func IsEmptyBytes(v []byte) bool {
	return len(v) <= 0
}

func IsNotEmptyBytes(v []byte) bool {
	return !IsEmptyBytes(v)
}

func IsEmptyInt(v []int) bool {
	return len(v) <= 0
}

func IsNotEmptyInt(v []int) bool {
	return !IsEmptyInt(v)
}

func IsEmptyInt32(v []int32) bool {
	return len(v) <= 0
}

func IsNotEmptyInt32(v []int32) bool {
	return !IsEmptyInt32(v)
}

func IsEmptyInt64(v []int64) bool {
	return len(v) <= 0
}

func IsNotEmptyInt64(v []int64) bool {
	return !IsEmptyInt64(v)
}

func IsEmptyDouble(v []float64) bool {
	return len(v) <= 0
}

func IsNotEmptyDouble(v []float64) bool {
	return !IsEmptyDouble(v)
}
