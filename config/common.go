package config

var NullBytes = []byte("null")

const HomeShortCut = "~"


//常用时间 以下时间都是以秒为单位
const (
	OneSecond   = 1
	ThreeSecond = 3
	FiveSecond  = 5
	TenSecond   = 10
	HalfMinute = 30

	OneMinute   = 60 * OneSecond
	ThreeMinute = 3 * OneMinute
	FiveMinute  = 5 * OneMinute
	TenMinute   = 10 * OneMinute
	HalfHour    = 30 * OneMinute

	OneHour    = 60 * OneMinute
	SizHour    = 6 * OneHour
	TwelveHour = 12 * OneHour

	HalfDay = TwelveHour
	OneDay  = 24 * OneHour
)

const DefaultTimePattern = "2006-01-02T15:04:05.000Z0700"

const (
	DefaultPage = 1
	DefaultPageSize = 15
)