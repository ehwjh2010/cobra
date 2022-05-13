package time

import (
	"time"

	"github.com/ehwjh2010/viper/global"
)

//==========================Format===============================

func Time2Str(t time.Time, layout string, loc *time.Location) string {
	return t.In(loc).Format(layout)
}

func Time2LocalStrWithLay(t time.Time, layout string) string {
	return Time2Str(t, layout, time.Local)
}

func Time2UTCStrWithLay(t time.Time, layout string) string {
	return Time2Str(t, layout, time.UTC)
}

func Time2UTCStr(t time.Time) string {
	return Time2UTCStrWithLay(t, global.DefaultTimePattern)
}

func Time2LocalStr(t time.Time) string {
	return Time2LocalStrWithLay(t, global.DefaultTimePattern)
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

// Sec2UTCStrWithLay 秒级时间戳转UTC时区字符串
func Sec2UTCStrWithLay(sec int64, layout string) string {
	return RawSec2Str(sec, layout, time.UTC)
}

// Sec2UTCStr 秒级时间戳转UTC时区字符串
func Sec2UTCStr(sec int64) string {
	return Sec2UTCStrWithLay(sec, global.DefaultTimePattern)
}

// Sec2LocalStrWithLay 秒级时间戳转东八区字符串
func Sec2LocalStrWithLay(sec int64, layout string) string {
	return RawSec2Str(sec, layout, time.Local)
}

// Sec2LocalStr 秒级时间戳转东八区字符串
func Sec2LocalStr(sec int64) string {
	return Sec2LocalStrWithLay(sec, global.DefaultTimePattern)
}

// Str2TimeWithLay 字符串转time
func Str2TimeWithLay(str, layout string) (time.Time, error) {
	return time.Parse(layout, str)
}

// Str2Time 字符串转time
func Str2Time(str string) (time.Time, error) {
	return Str2TimeWithLay(str, global.DefaultTimePattern)
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

// MillSec2UTCStrWithLay 毫秒级时间戳转UTC时区字符串
func MillSec2UTCStrWithLay(msec int64, layout string) string {
	return RawMillSec2Str(msec, layout, time.UTC)
}

// MillSec2LocalStrWithLay 毫秒级时间戳转东八区字符串
func MillSec2LocalStrWithLay(msec int64, layout string) string {
	return RawMillSec2Str(msec, layout, time.Local)
}

// MillSec2UTCStr 毫秒级时间戳转UTC时区字符串
func MillSec2UTCStr(msec int64) string {
	return MillSec2UTCStrWithLay(msec, global.DefaultTimePattern)
}

// MillSec2LocalStr 毫秒级时间戳转东八区字符串
func MillSec2LocalStr(msec int64) string {
	return MillSec2LocalStrWithLay(msec, global.DefaultTimePattern)
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

// MicroSec2UTCStrWithLay 毫秒级时间戳转UTC时区字符串
func MicroSec2UTCStrWithLay(msec int64, layout string) string {
	return RawMicroSec2Str(msec, layout, time.UTC)
}

// MicroSec2LocalStrWithLay 毫秒级时间戳转东八区字符串
func MicroSec2LocalStrWithLay(msec int64, layout string) string {
	return RawMicroSec2Str(msec, layout, time.Local)
}

// MicroSec2UTCStr 毫秒级时间戳转UTC时区字符串
func MicroSec2UTCStr(msec int64) string {
	return MicroSec2UTCStrWithLay(msec, global.DefaultTimePattern)
}

// MicroSec2LocalStr 毫秒级时间戳转东八区字符串
func MicroSec2LocalStr(msec int64) string {
	return MicroSec2LocalStrWithLay(msec, global.DefaultTimePattern)
}
