package double

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAny2Double(t *testing.T) {
	Convey("Cast interface to double", t, func() {
		tests := []struct {
			Value   interface{}
			Success bool
		}{
			{nil, false},
			{"1", false},
			{0, false},
			{0.0, true},
			{0.22, true},
			{1.22, true},
			{"1.222", false},
			{true, false},
			{false, false},
		}

		for _, test := range tests {
			_, err := Any2Double(test.Value)
			So(err == nil, ShouldEqual, test.Success)
		}
	})
}

func TestMustAny2Double(t *testing.T) {
	Convey("Cast interface must to double", t, func() {
		tests := []struct {
			Value interface{}
			Dest  float64
		}{
			{0.22, 0.22},
			{1.22, 1.22},
		}

		for _, test := range tests {
			dst := MustAny2Double(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestStr2Double(t *testing.T) {
	Convey("Cast str to double", t, func() {
		tests := []struct {
			Value   string
			Success bool
		}{
			{"0.22", true},
			{"1.22", true},
			{"0", true},
			{"aaaa", false},
			{"nil", false},
		}

		for _, test := range tests {
			_, err := Str2Double(test.Value)
			So(err == nil, ShouldEqual, test.Success)
		}
	})
}

func TestMustStr2Double(t *testing.T) {
	Convey("Cast str must to double", t, func() {
		tests := []struct {
			Value string
			Dest  float64
		}{
			{"0.22", 0.22},
			{"1.22", 1.22},
			{"0", 0},
			{"0.0", 0},
		}

		for _, test := range tests {
			dst := MustStr2Double(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestDouble2Str(t *testing.T) {
	Convey("Cast double to str", t, func() {
		tests := []struct {
			Value float64
			Dest  string
		}{
			{0.22, "0.22"},
			{1.2244444444444, "1.2244444444444"},
			{0.0, "0"},
		}

		for _, test := range tests {
			dst := Double2Str(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestDouble2Any(t *testing.T) {
	Convey("Cast double to str", t, func() {
		foo := 2.01
		any := Double2Any(foo)
		So(any.(float64), ShouldEqual, foo)
	})
}

func TestDoubleSlice2Any(t *testing.T) {
	Convey("Cast double to str", t, func() {
		temp := []float64{1.2, 1, 5, 2.456}
		tmp := SliceDouble2Any(temp)
		slice2Double, err := SliceAny2Double(tmp)
		So(err, ShouldBeNil)
		So(SliceFloat64Equal(slice2Double, temp), ShouldBeTrue)

		temp = nil
		tmp = SliceDouble2Any(temp)
		So(tmp, ShouldBeEmpty)

		temp = make([]float64, 0)
		tmp = SliceDouble2Any(temp)
		slice2Double, err = SliceAny2Double(tmp)
		So(err, ShouldBeNil)
		So(SliceFloat64Equal(slice2Double, temp), ShouldBeTrue)

		anySlice2Double, err := SliceAny2Double(nil)
		So(err, ShouldBeNil)
		So(anySlice2Double, ShouldBeNil)

		foo := []interface{}{1.22, "1", true}
		double, err := SliceAny2Double(foo)
		So(err, ShouldNotBeNil)
		So(double, ShouldBeNil)

	})
}
