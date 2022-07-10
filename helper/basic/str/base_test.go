package str

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	Convey("Str basic method", t, func() {
		var a string
		So(IsEmpty(a), ShouldBeTrue)
		So(Size(a), ShouldEqual, 0)

		a = "art"
		So(IsNotEmpty(a), ShouldBeTrue)

		a = "张三and李四|王五"
		So(Size(a), ShouldEqual, 10)
	})
}

func TestIsNotEmptySlice(t *testing.T) {
	Convey("Str slice is empty", t, func() {
		var a []string
		So(IsEmptySlice(a), ShouldBeTrue)
		a = append(a, "a")
		So(IsNotEmptySlice(a), ShouldBeTrue)
	})
}

func TestSubStr(t *testing.T) {
	Convey("Str sub str", t, func() {
		s := "张三and李四|王五"
		So(SubStr(s, 0, 0), ShouldEqual, "")
		So(SubStr(s, 0, 1), ShouldEqual, "张")
		So(SubStr(s, 0, 2), ShouldEqual, "张三")
		So(SubStr(s, 3, 3), ShouldEqual, "")
		So(SubStr(s, 3, 4), ShouldEqual, "n")
		So(SubStr(s, 2, 1), ShouldEqual, "")
		So(SubStr(s, 1, 15), ShouldEqual, "三and李四|王五")
		So(SubStr(s, 10, 11), ShouldEqual, "")
	})
}

func TestSubStrWithCount(t *testing.T) {
	Convey("Str sub str with count", t, func() {
		s := "张三and李四|王五"
		So(SubStrWithCount(s, 0), ShouldEqual, "")
		So(SubStrWithCount(s, 1), ShouldEqual, "张")
		So(SubStrWithCount(s, 2), ShouldEqual, "张三")
		So(SubStrWithCount(s, 3), ShouldEqual, "张三a")
	})
}

func TestSubStrRevWithCount(t *testing.T) {
	Convey("Str sub str rev with count", t, func() {
		s := "张三and李四|王五"
		So(SubStrRevWithCount(s, 0), ShouldEqual, "")
		So(SubStrRevWithCount(s, 1), ShouldEqual, "五")
		So(SubStrRevWithCount(s, 2), ShouldEqual, "王五")
		So(SubStrRevWithCount(s, 3), ShouldEqual, "|王五")
	})
}
