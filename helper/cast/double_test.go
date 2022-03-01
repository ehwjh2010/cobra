package cast

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
