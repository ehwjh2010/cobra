package utils

func IsNil(v interface{}) bool {
	return v == nil
}

func IsNotNil(v interface{}) bool {
	return !IsNil(v)
}
