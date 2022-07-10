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
		So(IsNotEmptyAny(demo), ShouldBeTrue)

		demo2 := make([]interface{}, 0)
		So(IsEmptyAny(demo2), ShouldBeTrue)

		demo2 = append(demo2, 4)
		So(IsNotEmptyAny(demo), ShouldBeTrue)
	})
}

func TestIsEmptyBytesSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []byte
		So(IsEmptyBytes(demo), ShouldBeTrue)

		demo = append(demo, '1')
		So(IsNotEmptyBytes(demo), ShouldBeTrue)

		demo2 := make([]byte, 0)
		So(IsEmptyBytes(demo2), ShouldBeTrue)

		demo2 = append(demo2, '4')
		So(IsNotEmptyBytes(demo), ShouldBeTrue)
	})
}

func TestIsEmptyStrSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []string
		So(IsEmptyStr(demo), ShouldBeTrue)

		demo = append(demo, "!")
		So(IsNotEmptyStr(demo), ShouldBeTrue)

		demo2 := make([]string, 0)
		So(IsEmptyStr(demo2), ShouldBeTrue)

		demo2 = append(demo2, "3")
		So(IsNotEmptyStr(demo), ShouldBeTrue)
	})
}

func TestIsEmptyIntSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int
		So(IsEmptyInt(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsNotEmptyInt(demo), ShouldBeTrue)

		demo2 := make([]int, 0)
		So(IsEmptyInt(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsNotEmptyInt(demo), ShouldBeTrue)
	})
}

func TestIsEmptyInt32Slice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int32
		So(IsEmptyInt32(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsNotEmptyInt32(demo), ShouldBeTrue)

		demo2 := make([]int32, 0)
		So(IsEmptyInt32(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsNotEmptyInt32(demo), ShouldBeTrue)
	})
}

func TestIsEmptyInt64Slice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int64
		So(IsEmptyInt64(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsNotEmptyInt64(demo), ShouldBeTrue)

		demo2 := make([]int64, 0)
		So(IsEmptyInt64(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsNotEmptyInt64(demo), ShouldBeTrue)
	})
}

func TestIsEmptyDoubleSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []float64
		So(IsEmptyDouble(demo), ShouldBeTrue)

		demo = append(demo, 1.1)
		So(IsNotEmptyDouble(demo), ShouldBeTrue)

		demo2 := make([]float64, 0)
		So(IsEmptyDouble(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3.333)
		So(IsNotEmptyDouble(demo), ShouldBeTrue)
	})
}
