package enums

import "time"

const BJ = "Asia/Shanghai"

// 常用时间 以下时间都是以秒为单位
const (
	ZeroSecond    = 0
	OneSecond     = 1
	TwoSecond     = 2
	ThreeSecond   = 3
	FiveSecond    = 5
	TenSecond     = 10
	TwentySecond  = 20
	HalfOneMinute = 30

	OneMinute    = 60 * OneSecond
	TwoMinute    = 2 * OneMinute
	ThreeMinute  = 3 * OneMinute
	FiveMinute   = 5 * OneMinute
	TenMinute    = 10 * OneMinute
	TwentyMinute = 20 * OneMinute
	HalfOneHour  = 30 * OneMinute

	OneHour    = 60 * OneMinute
	TwoHour    = 2 * OneHour
	ThreeHour  = 3 * OneHour
	FiveHour   = 5 * OneHour
	SixHour    = 6 * OneHour
	TwelveHour = 12 * OneHour

	HalfOneDay = TwelveHour
	OneDay     = 24 * OneHour
)

// 常用时间 以下时间都是以time.Duration为单位
const (
	ZeroSecD    = time.Duration(0)
	OneSecD     = time.Duration(1) * time.Second
	ThreeSecD   = time.Duration(3) * time.Second
	FiveSecD    = time.Duration(5) * time.Second
	TenSecD     = time.Duration(10) * time.Second
	TwentySecD  = time.Duration(20) * time.Second
	HalfOneMinD = time.Duration(30) * time.Second

	OneMinD      = time.Duration(1) * time.Minute
	ThreeMinD    = time.Duration(3) * time.Minute
	FiveMinD     = time.Duration(5) * time.Minute
	TenMinD      = time.Duration(10) * time.Minute
	TwentyMinD   = time.Duration(20) * time.Minute
	HalfOneHourD = time.Duration(30) * time.Minute

	OneHourD    = time.Duration(1) * time.Hour
	TwoHourD    = time.Duration(2) * time.Hour
	ThreeHourD  = time.Duration(3) * time.Hour
	FiveHourD   = time.Duration(5) * time.Hour
	SixHourD    = time.Duration(6) * time.Hour
	HalfOneDayD = time.Duration(12) * time.Hour

	OneDayD = time.Duration(24) * time.Hour
)
