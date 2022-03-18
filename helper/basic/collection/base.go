package collection

func IsEmptyAnySlice(v []interface{}) bool {
	return len(v) <= 0
}

func IsEmptyStrSlice(v []string) bool {
	return len(v) <= 0
}

func IsEmptyBytesSlice(v []byte) bool {
	return len(v) <= 0
}

func IsEmptyIntSlice(v []int) bool {
	return len(v) <= 0
}

func IsEmptyInt32Slice(v []int32) bool {
	return len(v) <= 0
}

func IsEmptyInt64Slice(v []int64) bool {
	return len(v) <= 0
}

func IsEmptyDoubleSlice(v []float64) bool {
	return len(v) <= 0
}
