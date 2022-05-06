package collection

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsEmptyAnySlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []interface{}
		So(IsEmptyAny(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyAny(demo), ShouldBeFalse)

		demo2 := make([]interface{}, 0)
		So(IsEmptyAny(demo2), ShouldBeTrue)

		demo2 = append(demo2, 4)
		So(IsEmptyAny(demo), ShouldBeFalse)
	})
}

func TestIsEmptyBytesSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []byte
		So(IsEmptyBytes(demo), ShouldBeTrue)

		demo = append(demo, '1')
		So(IsEmptyBytes(demo), ShouldBeFalse)

		demo2 := make([]byte, 0)
		So(IsEmptyBytes(demo2), ShouldBeTrue)

		demo2 = append(demo2, '4')
		So(IsEmptyBytes(demo), ShouldBeFalse)
	})
}

func TestIsEmptyStrSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []string
		So(IsEmptyStr(demo), ShouldBeTrue)

		demo = append(demo, "!")
		So(IsEmptyStr(demo), ShouldBeFalse)

		demo2 := make([]string, 0)
		So(IsEmptyStr(demo2), ShouldBeTrue)

		demo2 = append(demo2, "3")
		So(IsEmptyStr(demo), ShouldBeFalse)
	})
}

func TestIsEmptyIntSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int
		So(IsEmptyInt(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyInt(demo), ShouldBeFalse)

		demo2 := make([]int, 0)
		So(IsEmptyInt(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsEmptyInt(demo), ShouldBeFalse)
	})
}

func TestIsEmptyInt32Slice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int32
		So(IsEmptyInt32(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyInt32(demo), ShouldBeFalse)

		demo2 := make([]int32, 0)
		So(IsEmptyInt32(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsEmptyInt32(demo), ShouldBeFalse)
	})
}

func TestIsEmptyInt64Slice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int64
		So(IsEmptyInt64(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyInt64(demo), ShouldBeFalse)

		demo2 := make([]int64, 0)
		So(IsEmptyInt64(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsEmptyInt64(demo), ShouldBeFalse)
	})
}

func TestIsEmptyDoubleSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []float64
		So(IsEmptyDouble(demo), ShouldBeTrue)

		demo = append(demo, 1.1)
		So(IsEmptyDouble(demo), ShouldBeFalse)

		demo2 := make([]float64, 0)
		So(IsEmptyDouble(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3.333)
		So(IsEmptyDouble(demo), ShouldBeFalse)
	})
}
