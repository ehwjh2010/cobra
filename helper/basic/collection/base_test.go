package collection

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsEmptyAnySlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []interface{}
		So(IsEmptyAnySlice(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyAnySlice(demo), ShouldBeFalse)

		demo2 := make([]interface{}, 0)
		So(IsEmptyAnySlice(demo2), ShouldBeTrue)

		demo2 = append(demo2, 4)
		So(IsEmptyAnySlice(demo), ShouldBeFalse)
	})
}

func TestIsEmptyBytesSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []byte
		So(IsEmptyBytesSlice(demo), ShouldBeTrue)

		demo = append(demo, '1')
		So(IsEmptyBytesSlice(demo), ShouldBeFalse)

		demo2 := make([]byte, 0)
		So(IsEmptyBytesSlice(demo2), ShouldBeTrue)

		demo2 = append(demo2, '4')
		So(IsEmptyBytesSlice(demo), ShouldBeFalse)
	})
}

func TestIsEmptyStrSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []string
		So(IsEmptyStrSlice(demo), ShouldBeTrue)

		demo = append(demo, "!")
		So(IsEmptyStrSlice(demo), ShouldBeFalse)

		demo2 := make([]string, 0)
		So(IsEmptyStrSlice(demo2), ShouldBeTrue)

		demo2 = append(demo2, "3")
		So(IsEmptyStrSlice(demo), ShouldBeFalse)
	})
}

func TestIsEmptyIntSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int
		So(IsEmptyIntSlice(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyIntSlice(demo), ShouldBeFalse)

		demo2 := make([]int, 0)
		So(IsEmptyIntSlice(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsEmptyIntSlice(demo), ShouldBeFalse)
	})
}

func TestIsEmptyInt32Slice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int32
		So(IsEmptyInt32Slice(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyInt32Slice(demo), ShouldBeFalse)

		demo2 := make([]int32, 0)
		So(IsEmptyInt32Slice(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsEmptyInt32Slice(demo), ShouldBeFalse)
	})
}

func TestIsEmptyInt64Slice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []int64
		So(IsEmptyInt64Slice(demo), ShouldBeTrue)

		demo = append(demo, 1)
		So(IsEmptyInt64Slice(demo), ShouldBeFalse)

		demo2 := make([]int64, 0)
		So(IsEmptyInt64Slice(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3)
		So(IsEmptyInt64Slice(demo), ShouldBeFalse)
	})
}

func TestIsEmptyDoubleSlice(t *testing.T) {
	Convey("test empty slice", t, func() {
		var demo []float64
		So(IsEmptyDoubleSlice(demo), ShouldBeTrue)

		demo = append(demo, 1.1)
		So(IsEmptyDoubleSlice(demo), ShouldBeFalse)

		demo2 := make([]float64, 0)
		So(IsEmptyDoubleSlice(demo2), ShouldBeTrue)

		demo2 = append(demo2, 3.333)
		So(IsEmptyDoubleSlice(demo), ShouldBeFalse)
	})
}
