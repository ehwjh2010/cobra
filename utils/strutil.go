package utils

import "unicode/utf8"

func IsEmptyStr(str string) bool {
	if len(str) == 0 {
		return true
	}

	return false
}

func StrSize(str string) int {
	return utf8.RuneCountInString(str)
}
