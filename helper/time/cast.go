package time

import (
	"github.com/ehwjh2010/viper/client/enums"
	"time"
)

//==========================Format===============================

func Time2StrWithLay(t time.Time, layout string) string {
	return t.Format(layout)
}

func Time2Str(t time.Time) string {
	return t.Format(enums.DefaultTimePattern)
}

//==========================Second===============================

// Sec2Time 秒级时间戳转时间对象
func Sec2Time(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// RawSec2Str 秒级时间戳转字符串
func RawSec2Str(sec int64, layout string, loc *time.Location) string {
	return Sec2Time(sec).In(loc).Format(layout)
}

// Sec2StrWithLay 秒级时间戳转UTC时区字符串
func Sec2StrWithLay(sec int64, layout string) string {
	return RawSec2Str(sec, layout, time.UTC)
}

// Sec2Str 秒级时间戳转UTC时区字符串
func Sec2Str(sec int64) string {
	return RawSec2Str(sec, enums.DefaultTimePattern, time.UTC)
}

// Sec2BjStrWithLay 秒级时间戳转东八区字符串
func Sec2BjStrWithLay(sec int64, layout string) string {
	loc := GetBJLocation()
	return RawSec2Str(sec, layout, loc)
}

// Sec2BjStr 秒级时间戳转东八区字符串
func Sec2BjStr(sec int64) string {
	loc := GetBJLocation()
	return RawSec2Str(sec, enums.DefaultTimePattern, loc)
}

func Str2TimeWithLay(str, layout string) (time.Time, error) {
	return time.Parse(layout, str)
}

func Str2Time(str string) (time.Time, error) {
	return time.Parse(enums.DefaultTimePattern, str)
}

//==========================MillSecond===========================

// MillSec2Time 毫秒级时间戳转时间对象
func MillSec2Time(msec int64) time.Time {
	return time.UnixMilli(msec)
}

// RawMillSec2Str 毫秒级时间戳转字符串
func RawMillSec2Str(sec int64, layout string, loc *time.Location) string {
	return MillSec2Time(sec).In(loc).Format(layout)
}

// MillSec2StrWithLay 毫秒级时间戳转UTC时区字符串
func MillSec2StrWithLay(msec int64, layout string) string {
	return RawMillSec2Str(msec, layout, time.UTC)
}

// MillSec2Str 毫秒级时间戳转UTC时区字符串
func MillSec2Str(msec int64) string {
	return RawMillSec2Str(msec, enums.DefaultTimePattern, time.UTC)
}

// MillSec2BjStrWithLay 毫秒级时间戳转东八区字符串
func MillSec2BjStrWithLay(msec int64, layout string) string {
	loc := GetBJLocation()
	return RawMillSec2Str(msec, layout, loc)
}

// MillSec2BjStr 毫秒级时间戳转东八区字符串
func MillSec2BjStr(msec int64) string {
	loc := GetBJLocation()
	return RawMillSec2Str(msec, enums.DefaultTimePattern, loc)
}

//==========================MicroSecond===========================

// MicroSec2Time 微秒级时间戳转时间对象
func MicroSec2Time(msec int64) time.Time {
	return time.UnixMicro(msec)
}

// RawMicroSec2Str 毫秒级时间戳转字符串
func RawMicroSec2Str(msec int64, layout string, loc *time.Location) string {
	return MicroSec2Time(msec).In(loc).Format(layout)
}

// MicroSec2StrWithLay 毫秒级时间戳转UTC时区字符串
func MicroSec2StrWithLay(msec int64, layout string) string {
	return RawMicroSec2Str(msec, layout, time.UTC)
}

// MicroSec2Str 毫秒级时间戳转UTC时区字符串
func MicroSec2Str(msec int64) string {
	return RawMicroSec2Str(msec, enums.DefaultTimePattern, time.UTC)
}

// MicroSec2BjStrWithLay 毫秒级时间戳转东八区字符串
func MicroSec2BjStrWithLay(msec int64, layout string) string {
	loc := GetBJLocation()
	return RawMicroSec2Str(msec, layout, loc)
}

// MicroSec2BjStr 毫秒级时间戳转东八区字符串
func MicroSec2BjStr(msec int64) string {
	loc := GetBJLocation()
	return RawMicroSec2Str(msec, enums.DefaultTimePattern, loc)
}
