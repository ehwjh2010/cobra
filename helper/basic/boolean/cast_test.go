package boolean

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAny2Bool(t *testing.T) {
	Convey("Cast interface to bool", t, func() {
		tests := []struct {
			Value   interface{}
			Success bool
		}{
			{"1", false},
			{"aaa", false},
			{"true", false},
			{"True", false},
			{true, true},
			{false, true},
		}

		for _, test := range tests {
			_, err := Any2Bool(test.Value)
			So(err == nil, ShouldEqual, test.Success)
		}
	})
}

func TestMustAny2Bool(t *testing.T) {
	Convey("Cast interface must to bool", t, func() {
		tests := []struct {
			Value interface{}
			Data  bool
		}{
			{true, true},
			{false, false},
		}

		for _, test := range tests {
			dst := MustAny2Bool(test.Value)
			So(dst, ShouldEqual, test.Data)
		}
	})
}

func TestStr2Bool(t *testing.T) {
	Convey("Cast str to bool", t, func() {
		tests := []struct {
			Value   string
			Success bool
		}{
			{"true", true},
			{"1", true},
			{"T", true},
			{"True", true},
			{"false", true},
			{"0", true},
			{"FALSE", true},
			{"False", true},
			{"aaa", false},
		}

		for _, test := range tests {
			_, err := Str2Bool(test.Value)
			So(err == nil, ShouldEqual, test.Success)
		}
	})
}

func TestMustStr2Bool(t *testing.T) {
	Convey("Cast str must to bool", t, func() {
		tests := []struct {
			Value   string
			Success bool
		}{
			{"true", true},
			{"1", true},
			{"T", true},
			{"True", true},
			{"false", false},
			{"0", false},
			{"FALSE", false},
			{"False", false},
			{"aaa", false},
		}

		for _, test := range tests {
			dst := MustStr2Bool(test.Value)
			So(dst, ShouldEqual, test.Success)
		}
	})
}

func TestBool2Str(t *testing.T) {
	Convey("Cast bool to str", t, func() {
		tests := []struct {
			Value   bool
			Success string
		}{
			{true, "true"},
			{false, "false"},
		}

		for _, test := range tests {
			dst := Bool2Str(test.Value)
			So(dst, ShouldEqual, test.Success)
		}
	})
}
