package random

import (
	"bytes"
	"github.com/ehwjh2010/viper/constant"
	"math/rand"
	"time"
)

func seed() {
	rand.Seed(time.Now().UnixNano())
}

var (
	asciiLettersLower    = []byte(constant.AsciiLowercase)
	asciiLettersLowerLen = len(asciiLettersLower)

	asciiLetters    = []byte(constant.AsciiLetters)
	asciiLettersLen = len(asciiLetters)

	asciiTotal    = []byte(constant.AsciiLetters + constant.Digits)
	asciiTotalLen = len(asciiTotal)
)

// Random 返回[0,1)的随机数
func Random() float64 {
	seed()
	return rand.Float64()
}

// RandInt 返回[0,n)的随机数
func RandInt(n int) int {
	seed()
	return rand.Intn(n)
}

// SimpleRandInt 返回[0,10)的随机数
func SimpleRandInt() int {
	seed()
	return rand.Intn(10)
}

func baseRandStrN(count int, source []byte, srcLen int) string {
	var buf bytes.Buffer

	for i := 0; i < count; i++ {
		buf.WriteByte(source[RandInt(srcLen)])
	}

	result := buf.String()
	return result
}

// RandAsciiStr 随机字符串, 源: abcdefghijklmnopqrstuvwxyz
func RandAsciiStr() string {
	count := 6
	return baseRandStrN(count, asciiLettersLower, asciiLettersLowerLen)
}

// RandAsciiStrN 随机字符串, 源: abcdefghijklmnopqrstuvwxyz
func RandAsciiStrN(count int) string {
	return baseRandStrN(count, asciiLettersLower, asciiLettersLowerLen)
}

// RandAsciiStrCase 随机字符串, 大小写敏感, 源: abcdefghijklmnopqrstuvwxyz + ABCDEFGHIJKLMNOPQRSTUVWXYZ
func RandAsciiStrCase() string {
	count := 6
	return baseRandStrN(count, asciiLetters, asciiLettersLen)
}

// RandAsciiStrNCase 随机字符串, 大小写敏感, 源: abcdefghijklmnopqrstuvwxyz + ABCDEFGHIJKLMNOPQRSTUVWXYZ
func RandAsciiStrNCase(count int) string {
	return baseRandStrN(count, asciiLetters, asciiLettersLen)
}

// RandStr 随机字符串, 大小写敏感, 源: abcdefghijklmnopqrstuvwxyz + ABCDEFGHIJKLMNOPQRSTUVWXYZ + 01234567
func RandStr() string {
	count := 6
	return baseRandStrN(count, asciiTotal, asciiTotalLen)
}

// RandStrN 随机字符串, 大小写敏感, 源: abcdefghijklmnopqrstuvwxyz + ABCDEFGHIJKLMNOPQRSTUVWXYZ + 01234567
func RandStrN(count int) string {
	return baseRandStrN(count, asciiTotal, asciiTotalLen)
}
