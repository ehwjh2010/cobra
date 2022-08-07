package collection

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGroupIntSlice(t *testing.T) {
	Convey("Group int slice", t, func() {
		slice := []int{1, 2, 3, 4, 5}
		intSlices := GroupIntSlice(slice, 5)
		So(intSlices, ShouldResemble, [][]int{{1, 2, 3, 4, 5}})

		intSlice2 := GroupIntSlice(slice, 3)
		So(intSlice2, ShouldResemble, [][]int{{1, 2, 3}, {4, 5}})

		intSlice3 := GroupIntSlice(slice, 7)
		So(intSlice3, ShouldResemble, [][]int{{1, 2, 3, 4, 5}})
	})
}

func TestGroupInt32Slice(t *testing.T) {
	Convey("Group int slice", t, func() {
		slice := []int32{1, 2, 3, 4, 5}
		intSlices := GroupInt32Slice(slice, 5)
		So(intSlices, ShouldResemble, [][]int32{{1, 2, 3, 4, 5}})

		intSlice2 := GroupInt32Slice(slice, 3)
		So(intSlice2, ShouldResemble, [][]int32{{1, 2, 3}, {4, 5}})

		intSlice3 := GroupInt32Slice(slice, 7)
		So(intSlice3, ShouldResemble, [][]int32{{1, 2, 3, 4, 5}})
	})
}

func TestGroupInt64Slice(t *testing.T) {
	Convey("Group int slice", t, func() {
		slice := []int64{1, 2, 3, 4, 5}
		intSlices := GroupInt64Slice(slice, 5)
		So(intSlices, ShouldResemble, [][]int64{{1, 2, 3, 4, 5}})

		intSlice2 := GroupInt64Slice(slice, 3)
		So(intSlice2, ShouldResemble, [][]int64{{1, 2, 3}, {4, 5}})

		intSlice3 := GroupInt64Slice(slice, 7)
		So(intSlice3, ShouldResemble, [][]int64{{1, 2, 3, 4, 5}})
	})
}

func TestGroupStrSlice(t *testing.T) {
	Convey("Group int slice", t, func() {
		slice := []string{"1", "2", "3", "4", "5"}
		intSlices := GroupStrSlice(slice, 5)
		So(intSlices, ShouldResemble, [][]string{{"1", "2", "3", "4", "5"}})

		intSlice2 := GroupStrSlice(slice, 3)
		So(intSlice2, ShouldResemble, [][]string{{"1", "2", "3"}, {"4", "5"}})

		intSlice3 := GroupStrSlice(slice, 7)
		So(intSlice3, ShouldResemble, [][]string{{"1", "2", "3", "4", "5"}})
	})
}

func TestGroupFloat32Slice(t *testing.T) {
	Convey("Group int slice", t, func() {
		slice := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
		intSlices := GroupFloat32Slice(slice, 5)
		So(intSlices, ShouldResemble, [][]float32{{1.1, 2.2, 3.3, 4.4, 5.5}})

		intSlice2 := GroupFloat32Slice(slice, 3)
		So(intSlice2, ShouldResemble, [][]float32{{1.1, 2.2, 3.3}, {4.4, 5.5}})

		intSlice3 := GroupFloat32Slice(slice, 7)
		So(intSlice3, ShouldResemble, [][]float32{{1.1, 2.2, 3.3, 4.4, 5.5}})
	})
}

func TestGroupFloat64Slice(t *testing.T) {
	Convey("Group int slice", t, func() {
		slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
		intSlices := GroupFloat64Slice(slice, 5)
		So(intSlices, ShouldResemble, [][]float64{{1.1, 2.2, 3.3, 4.4, 5.5}})

		intSlice2 := GroupFloat64Slice(slice, 3)
		So(intSlice2, ShouldResemble, [][]float64{{1.1, 2.2, 3.3}, {4.4, 5.5}})

		intSlice3 := GroupFloat64Slice(slice, 7)
		So(intSlice3, ShouldResemble, [][]float64{{1.1, 2.2, 3.3, 4.4, 5.5}})
	})
}

func TestGroupAnySlice(t *testing.T) {
	Convey("Group int slice", t, func() {
		slice := []interface{}{1.1, 2.2, 3.3, 4.4, 5.5}
		intSlices := GroupAnySlice(slice, 5)
		So(intSlices, ShouldResemble, [][]interface{}{{1.1, 2.2, 3.3, 4.4, 5.5}})

		intSlice2 := GroupAnySlice(slice, 3)
		So(intSlice2, ShouldResemble, [][]interface{}{{1.1, 2.2, 3.3}, {4.4, 5.5}})

		intSlice3 := GroupAnySlice(slice, 7)
		So(intSlice3, ShouldResemble, [][]interface{}{{1.1, 2.2, 3.3, 4.4, 5.5}})
	})
}
