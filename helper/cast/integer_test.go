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
			Dest    int32
		}{
			{int32(math.MaxInt32), true, math.MaxInt32},
			{int32(math.MinInt32), true, math.MinInt32},
			{int32(0), true, 0},
			{"aaaa", false, 0},
		}

		for _, test := range tests {
			dst, err := Any2Int32(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
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
			Dest    int64
		}{
			{int64(math.MaxInt64), true, math.MaxInt64},
			{int64(math.MinInt64), true, math.MinInt64},
			{int64(0), true, 0},
			{2.22, false, 0},
			{2, false, 0},
			{"aaaaaa", false, 0},
			{true, false, 0},
		}

		for _, test := range tests {
			dst, err := Any2Int64(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustAny2Int64(t *testing.T) {
	Convey("Cast interface to int64", t, func() {
		tests := []struct {
			Value interface{}
			Dest  int64
		}{
			{int64(math.MaxInt64), math.MaxInt64},
			{int64(math.MinInt64), math.MinInt64},
			{int64(0), 0},
		}

		for _, test := range tests {
			dst := MustAny2Int64(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestAny2Int(t *testing.T) {
	Convey("Cast interface to int", t, func() {
		tests := []struct {
			Value   interface{}
			Success bool
			Dest    int
		}{
			{math.MaxInt, true, math.MaxInt},
			{math.MinInt, true, math.MinInt},
			{0, true, 0},
			{2.22, false, 0},
			{2, true, 2},
			{"aaaaaa", false, 0},
			{true, false, 0},
		}

		for _, test := range tests {
			dst, err := Any2Int(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustAny2Int(t *testing.T) {
	Convey("Cast interface to int", t, func() {
		tests := []struct {
			Value interface{}
			Dest  int
		}{
			{math.MaxInt64, math.MaxInt64},
			{math.MinInt64, math.MinInt64},
			{0, 0},
		}

		for _, test := range tests {
			dst := MustAny2Int(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestStr2Int(t *testing.T) {
	Convey("Cast str to int", t, func() {
		tests := []struct {
			Value   string
			Success bool
			Dest    int
		}{
			{"1111", true, 1111},
			{"aaaaa", false, 0},
			{"0", true, 0},
			{"0.0", false, 0},
			{"true", false, 0},
		}

		for _, test := range tests {
			dst, err := Str2Int(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustStr2Int(t *testing.T) {
	Convey("Cast str must to int", t, func() {
		tests := []struct {
			Value string
			Dest  int
		}{
			{"1111", 1111},
			{"0", 0},
		}

		for _, test := range tests {
			dst := MustStr2Int(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestStr2Int32(t *testing.T) {
	Convey("Cast str to int32", t, func() {
		tests := []struct {
			Value   string
			Success bool
			Dest    int32
		}{
			{"1111", true, 1111},
			{"aaaaa", false, 0},
			{"0", true, 0},
			{"0.0", false, 0},
			{"true", false, 0},
		}

		for _, test := range tests {
			dst, err := Str2Int32(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustStr2Int32(t *testing.T) {
	Convey("Cast str must to int32", t, func() {
		tests := []struct {
			Value string
			Dest  int32
		}{
			{"1111", 1111},
			{"0", 0},
		}

		for _, test := range tests {
			dst := MustStr2Int32(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestStr2Int64(t *testing.T) {
	Convey("Cast str to int64", t, func() {
		tests := []struct {
			Value   string
			Success bool
			Dest    int64
		}{
			{"1111", true, 1111},
			{"aaaaa", false, 0},
			{"0", true, 0},
			{"0.0", false, 0},
			{"true", false, 0},
		}

		for _, test := range tests {
			dst, err := Str2Int64(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustStr2Int64(t *testing.T) {
	Convey("Cast str must to int32", t, func() {
		tests := []struct {
			Value string
			Dest  int64
		}{
			{"1111", 1111},
			{"0", 0},
		}

		for _, test := range tests {
			dst := MustStr2Int64(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}
