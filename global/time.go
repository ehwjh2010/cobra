package global

import "time"

const BJ = "Asia/Shanghai"

const DefaultTimePattern = time.RFC3339 //默认时间格式

//常用时间 以下时间都是以秒为单位
const (
	OneSecond   = 1
	TwoSecond   = 2
	ThreeSecond = 3
	FiveSecond  = 5
	TenSecond   = 10
	HalfMinute  = 30

	OneMinute   = 60 * OneSecond
	TwoMinute   = 2 * OneMinute
	ThreeMinute = 3 * OneMinute
	FiveMinute  = 5 * OneMinute
	TenMinute   = 10 * OneMinute
	HalfHour    = 30 * OneMinute

	OneHour    = 60 * OneMinute
	TwoHour    = 2 * OneHour
	SixHour    = 6 * OneHour
	TwelveHour = 12 * OneHour

	HalfDay = TwelveHour
	OneDay  = 24 * OneHour
)

const (
	OneSecDur   = time.Duration(1) * time.Second
	ThreeSecDur = time.Duration(3) * time.Second
	FiveSecDur  = time.Duration(5) * time.Second
	TenSecDur   = time.Duration(10) * time.Second
	HalfMinDur  = time.Duration(30) * time.Second

	OneMinDur   = time.Duration(1) * time.Minute
	ThreeMinDur = time.Duration(3) * time.Minute
	FiveMinDur  = time.Duration(5) * time.Minute
	TenMinDur   = time.Duration(10) * time.Minute
	HalfHourDur = time.Duration(30) * time.Minute

	OneHourDur = time.Duration(1) * time.Hour
	SixHourDur = time.Duration(6) * time.Hour
	HalfDayDur = time.Duration(12) * time.Hour

	OneDayDur = time.Duration(24) * time.Hour
)
