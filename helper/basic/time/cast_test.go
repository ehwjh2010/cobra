package time

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTime2LocalStr(t *testing.T) {
	Convey("time to local str", t, func() {
		tests := []struct {
			Value time.Time
			Dest  string
			OK    bool
		}{
			{time.Unix(1645190888, 0), "2022-02-18T21:28:08+08:00", true},
			{time.Unix(1645190888, 0), "2022-02-18T13:28:08Z", false},
		}

		var ret string
		for _, test := range tests {
			ret = Time2LocalStr(test.Value)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}

func TestTime2UTCStr(t *testing.T) {
	Convey("time to utc str", t, func() {
		tests := []struct {
			Value time.Time
			Dest  string
			OK    bool
		}{
			{time.Unix(1645190888, 0), "2022-02-18T21:28:08+08:00", false},
			{time.Unix(1645190888, 0), "2022-02-18T13:28:08Z", true},
		}

		var ret string
		for _, test := range tests {
			ret = Time2UTCStr(test.Value)
			fmt.Println(ret)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}

func TestSec2UTCStr(t *testing.T) {
	Convey("stamp to utc str", t, func() {
		tests := []struct {
			Value int64
			Dest  string
			OK    bool
		}{
			{1645190888, "2022-02-18T21:28:08+08:00", false},
			{1645190888, "2022-02-18T13:28:08Z", true},
		}

		var ret string
		for _, test := range tests {
			ret = Sec2UTCStr(test.Value)
			fmt.Println(ret)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}

func TestSec2LocalStr(t *testing.T) {
	Convey("stamp to local str", t, func() {
		tests := []struct {
			Value int64
			Dest  string
			OK    bool
		}{
			{1645190888, "2022-02-18T21:28:08+08:00", true},
			{1645190888, "2022-02-18T13:28:08Z", false},
		}

		var ret string
		for _, test := range tests {
			ret = Sec2LocalStr(test.Value)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}

func TestStr2Time(t *testing.T) {
	Convey("stamp to local str", t, func() {
		tests := []struct {
			Value string
			Dest  time.Time
			OK    bool
		}{
			{"2022-02-18T13:28:08Z", time.Unix(1645190888, 0), true},
			{"2022-02-18T21:28:08+08:00", time.Unix(1645190888, 0), true},
			{"aaaaaaa", time.Unix(1645190888, 0), false},
		}

		var ret time.Time
		var err error
		for _, test := range tests {
			ret, err = Str2Time(test.Value)
			So(err == nil, ShouldEqual, test.OK)
			if test.OK {
				So(ret.Equal(test.Dest), ShouldBeTrue)
			}
		}
	})
}

func TestMillSec2UTCStr(t *testing.T) {
	Convey("stamp to utc str", t, func() {
		tests := []struct {
			Value int64
			Dest  string
			OK    bool
		}{
			{1645190888000, "2022-02-18T21:28:08+08:00", false},
			{1645190888000, "2022-02-18T13:28:08Z", true},
		}

		var ret string
		for _, test := range tests {
			ret = MillSec2UTCStr(test.Value)
			fmt.Println(ret)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}

func TestMillSec2LocalStr(t *testing.T) {
	Convey("stamp to local str", t, func() {
		tests := []struct {
			Value int64
			Dest  string
			OK    bool
		}{
			{1645190888000, "2022-02-18T21:28:08+08:00", true},
			{1645190888000, "2022-02-18T13:28:08Z", false},
		}

		var ret string
		for _, test := range tests {
			ret = MillSec2LocalStr(test.Value)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}

func TestMicroSec2UTCStr(t *testing.T) {
	Convey("stamp to utc str", t, func() {
		tests := []struct {
			Value int64
			Dest  string
			OK    bool
		}{
			{1645190888000000, "2022-02-18T21:28:08+08:00", false},
			{1645190888000000, "2022-02-18T13:28:08Z", true},
		}

		var ret string
		for _, test := range tests {
			ret = MicroSec2UTCStr(test.Value)
			fmt.Println(ret)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}

func TestMicroSec2LocalStr(t *testing.T) {
	Convey("stamp to local str", t, func() {
		tests := []struct {
			Value int64
			Dest  string
			OK    bool
		}{
			{1645190888000000, "2022-02-18T21:28:08+08:00", true},
			{1645190888000000, "2022-02-18T13:28:08Z", false},
		}

		var ret string
		for _, test := range tests {
			ret = MicroSec2LocalStr(test.Value)
			So(ret == test.Dest, ShouldEqual, test.OK)
		}
	})
}
