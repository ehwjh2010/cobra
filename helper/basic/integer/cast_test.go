package integer

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
			{int32(0), 0},
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

func TestInteger2Any(t *testing.T) {
	Convey("Cast integer to interface", t, func() {
		foo := math.MaxInt
		any := Int2Any(foo)
		any2Int, err := Any2Int(any)
		So(err, ShouldBeNil)
		So(any2Int, ShouldEqual, foo)

		foo = math.MinInt
		any = Int2Any(foo)
		any2Int, err = Any2Int(any)
		So(err, ShouldBeNil)
		So(any2Int, ShouldEqual, foo)
	})
}

func TestInt32TAny(t *testing.T) {
	Convey("Cast int32 slice to any", t, func() {
		var a int32

		any := Int32TAny(a)
		So(any, ShouldNotBeNil)

	})
}

func TestInt2Str(t *testing.T) {
	Convey("Cast int to str", t, func() {
		var a = 30
		r := Int2Str(a)
		So(r, ShouldEqual, "30")

	})
}

func TestInt32ToStr(t *testing.T) {
	Convey("Cast int to str", t, func() {
		var a int32 = 30
		r := Int32ToStr(a)
		So(r, ShouldEqual, "30")

	})
}

func TestInt64ToStr(t *testing.T) {
	Convey("Cast int to str", t, func() {
		var a int64 = 30
		r := Int64ToStr(a)
		So(r, ShouldEqual, "30")

	})
}

func TestInt64ToAny(t *testing.T) {
	Convey("Cast int to any", t, func() {
		var a int64 = 30
		any := Int64ToAny(a)
		So(any, ShouldEqual, interface{}(a))
	})
}
