package strutils

import "unicode/utf8"

func IsEmptyStr(str string) bool {
	if len(str) == 0 {
		return true
	}

	return false
}

func IsNotEmptyStr(str string) bool {
	return !IsEmptyStr(str)
}

func StrSize(str string) int {
	return utf8.RuneCountInString(str)
}
