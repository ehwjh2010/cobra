package collection

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInt32Slice2Any(t *testing.T) {
	convey.Convey("Cast integer slice to interface", t, func() {
		foo := []int32{1, 2, 3, 4}
		any := Int32Slice2AnySlice(foo)
		any2Int, err := AnySlice2Int32Slice(any)
		convey.So(err, convey.ShouldBeNil)
		convey.So(any2Int, convey.ShouldResemble, foo)

		slice2Any := Int32Slice2AnySlice(nil)
		convey.So(slice2Any, convey.ShouldBeEmpty)

		temp := []interface{}{true, 1, 3}
		any2Int, err = AnySlice2Int32Slice(temp)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(any2Int, convey.ShouldBeNil)

		slice2Int32, err := AnySlice2Int32Slice(nil)
		convey.So(err, convey.ShouldBeNil)
		convey.So(slice2Int32, convey.ShouldBeNil)
	})
}

func TestSliceInt64TAny(t *testing.T) {
	convey.Convey("Cast integer slice to interface", t, func() {
		foo := []int64{1, 2, 3, 4}
		any := Int64Slice2AnySlice(foo)
		any2Int, err := AnySlice2Int64Slice(any)
		convey.So(err, convey.ShouldBeNil)
		convey.So(any2Int, convey.ShouldResemble, foo)

		slice2Any := Int64Slice2AnySlice(nil)
		convey.So(slice2Any, convey.ShouldBeEmpty)

		temp := []interface{}{true, 1, 3}
		any2Int, err = AnySlice2Int64Slice(temp)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(any2Int, convey.ShouldBeNil)

		slice2Int64, err := AnySlice2Int64Slice(nil)
		convey.So(err, convey.ShouldBeNil)
		convey.So(slice2Int64, convey.ShouldBeEmpty)
	})
}

func TestMustAnySlice2Int(t *testing.T) {
	convey.Convey("Cast any slice to int slice", t, func() {
		foo := []int{1, 2, 3, 4}
		any := IntSlice2AnySlice(foo)
		ints := MustAnySlice2IntSlice(any)
		convey.So(ints, convey.ShouldResemble, foo)

		temp := make([]int, 0)
		any2 := IntSlice2AnySlice(temp)
		any2Int := MustAnySlice2IntSlice(any2)
		convey.So(any2Int, convey.ShouldResemble, temp)

		var bar []int
		any3 := IntSlice2AnySlice(bar)
		any3Int := MustAnySlice2IntSlice(any3)
		convey.So(any3Int, convey.ShouldResemble, bar)
	})
}

func TestMustAnySlice2Int32(t *testing.T) {
	convey.Convey("Cast any slice to int32 slice", t, func() {
		foo := []int32{1, 2, 3, 4}
		any := Int32Slice2AnySlice(foo)
		ints := MustAnySlice2Int32(any)
		convey.So(ints, convey.ShouldResemble, foo)

		temp := make([]int32, 0)
		any2 := Int32Slice2AnySlice(temp)
		any2Int := MustAnySlice2Int32(any2)
		convey.So(any2Int, convey.ShouldResemble, temp)

		var bar []int32
		any3 := Int32Slice2AnySlice(bar)
		any3Int := MustAnySlice2Int32(any3)
		convey.So(any3Int, convey.ShouldResemble, bar)
	})
}

func TestMustAnySlice2Int64(t *testing.T) {
	convey.Convey("Cast any slice to int64 slice", t, func() {
		foo := []int64{1, 2, 3, 4}
		any := Int64Slice2AnySlice(foo)
		ints := MustAnySlice2Int64Slice(any)
		convey.So(ints, convey.ShouldResemble, foo)

		temp := make([]int64, 0)
		any2 := Int64Slice2AnySlice(temp)
		any2Int := MustAnySlice2Int64Slice(any2)
		convey.So(any2Int, convey.ShouldResemble, temp)

		var bar []int64
		any3 := Int64Slice2AnySlice(bar)
		any3Int := MustAnySlice2Int64Slice(any3)
		convey.So(any3Int, convey.ShouldResemble, bar)
	})
}

func TestSliceInt2Int32(t *testing.T) {
	convey.Convey("Cast int slice to int32 slice", t, func() {
		foo := []int{1, 2, 3, 4}
		ints := IntSlice2Int32Slice(foo)
		dest := []int32{1, 2, 3, 4}
		convey.So(ints, convey.ShouldResemble, dest)

		foo = make([]int, 0)
		ints = IntSlice2Int32Slice(foo)
		dest = make([]int32, 0)
		convey.So(ints, convey.ShouldResemble, dest)

		var temp []int
		ints = IntSlice2Int32Slice(temp)
		convey.So(ints, convey.ShouldEqual, nil)

	})
}

func TestInt32Slice2IntSlice(t *testing.T) {
	convey.Convey("Cast int32 slice to int slice", t, func() {
		foo := []int32{1, 2, 3, 4}
		ints := Int32Slice2IntSlice(foo)
		dest := []int{1, 2, 3, 4}
		convey.So(ints, convey.ShouldResemble, dest)

		foo = make([]int32, 0)
		ints = Int32Slice2IntSlice(foo)
		dest = make([]int, 0)
		convey.So(ints, convey.ShouldResemble, dest)

		var temp []int32
		ints = Int32Slice2IntSlice(temp)
		convey.So(ints, convey.ShouldEqual, nil)

	})
}

func TestIntSlice2Int64Slice(t *testing.T) {
	convey.Convey("Cast int slice to int64 slice", t, func() {
		foo := []int{1, 2, 3, 4}
		ints := IntSlice2Int64Slice(foo)
		dest := []int64{1, 2, 3, 4}
		convey.So(ints, convey.ShouldResemble, dest)

		foo = make([]int, 0)
		ints = IntSlice2Int64Slice(foo)
		dest = make([]int64, 0)
		convey.So(ints, convey.ShouldResemble, dest)

		var temp []int
		ints = IntSlice2Int64Slice(temp)
		convey.So(ints, convey.ShouldEqual, nil)

	})
}

func TestInt64Slice2IntSlice(t *testing.T) {
	convey.Convey("Cast int64 slice to int slice", t, func() {
		foo := []int64{1, 2, 3, 4}
		ints := Int64Slice2IntSlice(foo)
		dest := []int{1, 2, 3, 4}
		convey.So(ints, convey.ShouldResemble, dest)

		foo = make([]int64, 0)
		ints = Int64Slice2IntSlice(foo)
		dest = make([]int, 0)
		convey.So(ints, convey.ShouldResemble, dest)

		var temp []int64
		ints = Int64Slice2IntSlice(temp)
		convey.So(ints, convey.ShouldEqual, nil)

	})
}

func TestInt32Slice2Int64Slice(t *testing.T) {
	convey.Convey("Cast int32 slice to int64 slice", t, func() {
		foo := []int32{1, 2, 3, 4}
		ints := Int32Slice2Int64Slice(foo)
		dest := []int64{1, 2, 3, 4}
		convey.So(ints, convey.ShouldResemble, dest)

		foo = make([]int32, 0)
		ints = Int32Slice2Int64Slice(foo)
		dest = make([]int64, 0)
		convey.So(ints, convey.ShouldResemble, dest)

		var temp []int32
		ints = Int32Slice2Int64Slice(temp)
		convey.So(ints, convey.ShouldEqual, nil)

	})
}

func TestInt64Slice2Int32Slice(t *testing.T) {
	convey.Convey("Cast int64 slice to int32 slice", t, func() {
		foo := []int64{1, 2, 3, 4}
		ints := Int64Slice2Int32Slice(foo)
		dest := []int32{1, 2, 3, 4}
		convey.So(ints, convey.ShouldResemble, dest)

		foo = make([]int64, 0)
		ints = Int64Slice2Int32Slice(foo)
		dest = make([]int32, 0)
		convey.So(ints, convey.ShouldResemble, dest)

		var temp []int64
		ints = Int64Slice2Int32Slice(temp)
		convey.So(ints, convey.ShouldEqual, nil)

	})
}

func TestIntSlice2Any(t *testing.T) {
	convey.Convey("Cast integer slice to interface", t, func() {
		foo := []int{1, 2, 3, 4}
		any := IntSlice2AnySlice(foo)
		any2Int, err := AnySlice2IntSlice(any)
		convey.So(err, convey.ShouldBeNil)
		convey.So(any2Int, convey.ShouldResemble, foo)

		slice2Any := IntSlice2AnySlice(nil)
		convey.So(slice2Any, convey.ShouldBeEmpty)

		temp := []interface{}{true, 1, 3}
		any2Int, err = AnySlice2IntSlice(temp)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(any2Int, convey.ShouldBeNil)

		slice2Int, err := AnySlice2IntSlice(nil)
		convey.So(err, convey.ShouldBeNil)
		convey.So(slice2Int, convey.ShouldBeNil)
	})
}
