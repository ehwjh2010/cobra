package str

import (
	"unicode/utf8"
)

func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func Size(str string) int {
	return utf8.RuneCountInString(str)
}
