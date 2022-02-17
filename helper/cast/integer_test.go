package cast

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestInt32ToInt(t *testing.T) {
	Convey("Cast int32 to int", t, func() {
		tests := []struct {
			Value int32
			Dest  int
		}{
			{math.MaxInt32, math.MaxInt32},
			{math.MinInt32, math.MinInt32},
			{1, 1},
		}

		for _, test := range tests {
			dst := Int32ToInt(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestInt64ToInt(t *testing.T) {
	Convey("Cast int64 to int", t, func() {
		tests := []struct {
			Value int64
			Dest  int
		}{
			// 如果机器是32位, 下面的case会溢出
			{math.MaxInt64, math.MaxInt64},
			{math.MinInt64, math.MinInt64},
			{1, 1},
		}

		for _, test := range tests {
			dst := Int64ToInt(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestIntToInt32(t *testing.T) {
	Convey("Cast int to int32", t, func() {
		tests := []struct {
			Value int
			Dest  int32
		}{
			{math.MaxInt32, math.MaxInt32},
			{math.MinInt32, math.MinInt32},
			{1, 1},
		}

		for _, test := range tests {
			dst := IntToInt32(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestInt64ToInt32(t *testing.T) {
	Convey("Cast int64 to int32", t, func() {
		tests := []struct {
			Value int64
			Dest  int32
		}{
			{math.MaxInt32, math.MaxInt32},
			{math.MinInt32, math.MinInt32},
			{1, 1},
			// 以下会溢出
			//{math.MaxInt64, math.MaxInt64},
			//{math.MinInt64, math.MinInt64},
		}

		for _, test := range tests {
			dst := Int64ToInt32(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestIntToInt64(t *testing.T) {
	Convey("Cast int to int64", t, func() {
		tests := []struct {
			Value int
			Dest  int64
		}{
			{math.MaxInt, math.MaxInt},
			{math.MinInt, math.MinInt},
			{1, 1},
		}

		for _, test := range tests {
			dst := IntToInt64(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestInt32ToInt64(t *testing.T) {
	Convey("Cast int32 to int64", t, func() {
		tests := []struct {
			Value int32
			Dest  int64
		}{
			{math.MaxInt32, math.MaxInt32},
			{math.MinInt32, math.MinInt32},
			{1, 1},
		}

		for _, test := range tests {
			dst := Int32ToInt64(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestAny2Int32(t *testing.T) {
	Convey("Cast interface to int32", t, func() {
		tests := []struct {
			Value   interface{}
			Success bool
		}{
			{int32(math.MaxInt32), true},
			{int32(math.MinInt32), true},
			{int32(0), true},
			{"aaaa", false},
		}

		for _, test := range tests {
			_, err := Any2Int32(test.Value)
			So(err == nil, ShouldEqual, test.Success)
		}
	})
}

func TestMustAny2Int32(t *testing.T) {
	Convey("Cast interface must to int32", t, func() {
		tests := []struct {
			Value interface{}
			Dest  int32
		}{
			{int32(math.MaxInt32), math.MaxInt32},
			{int32(math.MinInt32), math.MinInt32},
			{0, 0},
		}

		for _, test := range tests {
			dst := MustAny2Int32(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestAny2Int64(t *testing.T) {
	Convey("Cast interface to int64", t, func() {
		tests := []struct {
			Value   interface{}
			Success bool
		}{
			{math.MaxInt64, true},
			{math.MinInt64, true},
			{0, true},
			{"aaaaaa", false},
			{true, false},
		}

		for _, test := range tests {
			_, err := Any2Int64(test.Value)
			So(err == nil, ShouldEqual, test.Success)
		}
	})
}
