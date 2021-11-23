package timeutils

import (
	"github.com/ehwjh2010/cobra/config"
	"time"
)

func Time2StrWithLay(t time.Time, layout string) string {
	return t.Format(layout)
}

func Time2Str(t time.Time) string {
	return t.Format(config.DefaultTimePattern)
}

func Str2Time(str string) (time.Time, error) {
	return time.Parse(config.DefaultTimePattern, str)
}

func Str2TimeWithLay(str string, layout string) (time.Time, error) {
	return time.Parse(layout, str)
}