package str

import (
	"unicode/utf8"
)

func IsEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}

	return false
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func Size(str string) int {
	return utf8.RuneCountInString(str)
}
