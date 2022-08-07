package collection

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCharSlice2AnySlice(t *testing.T) {
	convey.Convey("Cast slice string to any", t, func() {
		foo := []string{"a", "b", "c", "d"}
		temp := CharSlice2AnySlice(foo)
		slice2Str, err := AnySlice2CharSlice(temp)
		convey.So(err, convey.ShouldBeNil)
		convey.So(slice2Str, convey.ShouldResemble, foo)

		v := CharSlice2AnySlice(nil)
		convey.So(v, convey.ShouldBeNil)

		tmp := make([]interface{}, 3)
		tmp[0] = 1
		tmp[0] = "1"
		tmp[0] = true

		slice2Str, err = AnySlice2CharSlice(tmp)
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(slice2Str, convey.ShouldBeEmpty)

		str, err := AnySlice2CharSlice(nil)
		convey.So(err, convey.ShouldBeNil)
		convey.So(str, convey.ShouldBeEmpty)
	})
}

func TestMustAnySlice2CharSlice(t *testing.T) {
	convey.Convey("Cast must any slice to string slice", t, func() {
		foo := []interface{}{"a", "b", "c", "d"}
		temp := MustAnySlice2CharSlice(foo)
		dest := []string{"a", "b", "c", "d"}
		convey.So(temp, convey.ShouldResemble, dest)

		v := MustAnySlice2CharSlice(nil)
		convey.So(v, convey.ShouldBeNil)
	})
}

func TestAnySlice2CharSlice(t *testing.T) {
	convey.Convey("Cast any slice to string slice", t, func() {
		foo := []interface{}{"a", "b", "c", "d"}
		temp, err := AnySlice2CharSlice(foo)
		dest := []string{"a", "b", "c", "d"}
		convey.So(err, convey.ShouldBeNil)
		convey.So(temp, convey.ShouldResemble, dest)

		v := CharSlice2AnySlice(nil)
		convey.So(v, convey.ShouldBeNil)
	})
}
