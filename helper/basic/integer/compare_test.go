package integer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMinInt(t *testing.T) {
	Convey("compare min int value", t, func() {
		minInt := MinInt(-1, 2, 3, 0, 9)
		So(minInt, ShouldEqual, -1)
	})
}

func TestMinInt32(t *testing.T) {
	Convey("compare min int32 value", t, func() {
		minInt32 := MinInt32(-1, 2, 3, 0, 9)
		So(minInt32, ShouldEqual, -1)
	})
}

func TestMinInt64(t *testing.T) {
	Convey("compare min int64 value", t, func() {
		minInt64 := MinInt64(-1, 2, 3, 0, 9)
		So(minInt64, ShouldEqual, -1)
	})
}

func TestMaxInt(t *testing.T) {
	Convey("compare max int value", t, func() {
		maxInt := MaxInt(-1, 2, 3, 0, 9)
		So(maxInt, ShouldEqual, 9)
	})
}

func TestMaxInt32(t *testing.T) {
	Convey("compare max int32 value", t, func() {
		maxInt32 := MaxInt32(-1, 2, 3, 0, 9)
		So(maxInt32, ShouldEqual, 9)
	})
}

func TestMaxInt64(t *testing.T) {
	Convey("compare max int64 value", t, func() {
		maxInt64 := MaxInt64(-1, 2, 3, 0, 9)
		So(maxInt64, ShouldEqual, 9)
	})
}
