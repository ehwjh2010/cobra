package collection

import (
	"github.com/ehwjh2010/viper/verror"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBoolSlice2AnySlice(t *testing.T) {
	convey.Convey("Cast bool to str", t, func() {
		var b []bool
		any := BoolSlice2AnySlice(b)
		convey.So(any, convey.ShouldBeEmpty)
		b = append(b, true)
		any = BoolSlice2AnySlice(b)
		convey.So(any, convey.ShouldNotBeEmpty)
	})
}

func TestAnySlice2BoolSlice(t *testing.T) {
	convey.Convey("Cast any slice to bool slice", t, func() {
		var b = []interface{}{true, false, true}
		any, err := AnySlice2BoolSlice(b)
		convey.So(err, convey.ShouldBeNil)
		convey.So(any, convey.ShouldResemble, []bool{true, false, true})

		b = []interface{}{1, 2, 3}
		_, err = AnySlice2BoolSlice(b)
		convey.So(err, convey.ShouldBeError, verror.CastBoolErr)
	})
}
