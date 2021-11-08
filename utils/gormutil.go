package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

const Layout = "2006-01-02 15:04:05"

type UTCTime time.Time

func (t *UTCTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(Layout, timeStr)
	*t = UTCTime(t1)
	return err
}

func (t UTCTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).In(time.UTC).Format(Layout))
	return []byte(formatted), nil
}

func (t UTCTime) Value() (driver.Value, error) {
	// UTCTime 转换成 time.Time 类型
	tTime := time.Time(t).In(time.UTC)
	return tTime.Format(Layout), nil
}

func (t *UTCTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = UTCTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *UTCTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).In(time.UTC).String())
}
