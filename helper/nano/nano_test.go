package nano

import (
	"github.com/ehwjh2010/viper/constant"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetNanoId(t *testing.T) {
	Convey("Get nano id", t, func() {
		_, err := GetNanoId()
		So(err, ShouldBeNil)
	})
}

func TestMustGetNanoId(t *testing.T) {
	Convey("Get nano id", t, func() {
		length := 21
		id := MustGetNanoId()
		So(len(id), ShouldEqual, length)
	})
}

func TestGetNonaIdWithCustom(t *testing.T) {
	Convey("Get nano id by custom", t, func() {
		size := 10
		id, err := GetNonaIdWithCustom(constant.AsciiLetters, size)
		So(err, ShouldBeNil)
		So(len(id), ShouldEqual, size)
	})
}
