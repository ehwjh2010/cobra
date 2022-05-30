package str

import "unicode/utf8"

func IsEmpty(str string) bool {
	return len(str) <= 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func Size(str string) int {
	return utf8.RuneCountInString(str)
}

func IsEmptySlice(v []string) bool {
	return len(v) <= 0
}

func IsNotEmptySlice(v []string) bool {
	return !IsEmptySlice(v)
}

// SubStr 字符串截取
func SubStr(s string, start, end int) string {
	if s == "" || start >= end {
		return ""
	}

	sLen := utf8.RuneCountInString(s)
	if start >= sLen {
		return ""
	}

	var size, n, st int
	for i := 0; i < end && n < len(s); i++ {
		if i == start {
			st = n
		}
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}
	return s[st:n]
}
