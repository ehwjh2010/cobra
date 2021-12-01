package config

import "time"

var NullBytes = []byte("null")

const SwaggerAPIUrl = "/swagger/index.html"

const BJ = "Asia/Shanghai"

const NullStr = "null"

const HomeShortCut = "~" //类Unix系统home路径的短符号

//常用时间 以下时间都是以秒为单位
const (
	OneSecond   = 1
	ThreeSecond = 3
	FiveSecond  = 5
	TenSecond   = 10
	HalfMinute  = 30

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

const DefaultTimePattern = time.RFC3339 //默认时间格式

const (
	DefaultPage     = 1  //默认页数
	DefaultPageSize = 15 //默认每页数据
)

const (
	CN = "cn" //中文
	EN = "en" //英文
)
