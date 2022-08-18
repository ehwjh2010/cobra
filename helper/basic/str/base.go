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

// SubStr 字符串截取, 遵循左闭右开原则.
func SubStr(s string, start, end int) string {
	if s == "" || start >= end || start < 0 || end < 0 {
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

// SubStrWithCount 字符串从前截取，前n个字符.
func SubStrWithCount(s string, count int) string {
	return SubStr(s, 0, count)
}

// SubStrRevWithCount 字符串从后截取，截取后n个字符.
func SubStrRevWithCount(s string, count int) string {
	size := Size(s)

	return SubStr(s, size-count, size)
}
