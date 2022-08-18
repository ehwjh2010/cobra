package str

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	Convey("Cast str to bytes", t, func() {
		tests := []struct {
			Value   string
			Dest    []byte
			Success bool
		}{
			{"aaa", []byte("aaa"), true},
			{"true", []byte("true"), true},
			{"2222", []byte("222"), false},
			{"12311", []byte("aaaaa"), false},
			{"", nil, true},
		}

		var ret []byte
		for _, test := range tests {
			ret = Char2Bytes(test.Value)
			So(SliceByteEqual(ret, test.Dest), ShouldEqual, test.Success)
		}
	})
}

func TestAny2String(t *testing.T) {
	Convey("Cast any to string", t, func() {
		var foo interface{}
		origin := "1"
		foo = origin

		any2String := MustAny2Char(foo)
		So(any2String, ShouldEqual, origin)

		foo = 2
		any2String, err := Any2Char(foo)
		So(err, ShouldNotBeNil)
		So(any2String, ShouldBeEmpty)
	})
}

func TestStr2Any(t *testing.T) {
	Convey("Cast string to any", t, func() {
		foo := "art"
		So(MustAny2Char(Char2Any(foo)), ShouldEqual, foo)
	})
}

func TestChar2Int(t *testing.T) {
	Convey("Cast str to int", t, func() {
		tests := []struct {
			Value   string
			Success bool
			Dest    int
		}{
			{"1111", true, 1111},
			{"aaaaa", false, 0},
			{"0", true, 0},
			{"0.0", false, 0},
			{"true", false, 0},
		}

		for _, test := range tests {
			dst, err := Char2Int(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustChar2Int(t *testing.T) {
	Convey("Cast str must to int", t, func() {
		tests := []struct {
			Value string
			Dest  int
		}{
			{"1111", 1111},
			{"0", 0},
		}

		for _, test := range tests {
			dst := MustChar2Int(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestChar2Int32(t *testing.T) {
	Convey("Cast str to int32", t, func() {
		tests := []struct {
			Value   string
			Success bool
			Dest    int32
		}{
			{"1111", true, 1111},
			{"aaaaa", false, 0},
			{"0", true, 0},
			{"0.0", false, 0},
			{"true", false, 0},
		}

		for _, test := range tests {
			dst, err := Char2Int32(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustChar2Int32(t *testing.T) {
	Convey("Cast str must to int32", t, func() {
		tests := []struct {
			Value string
			Dest  int32
		}{
			{"1111", 1111},
			{"0", 0},
		}

		for _, test := range tests {
			dst := MustChar2Int32(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}

func TestChar2Int64(t *testing.T) {
	Convey("Cast str to int64", t, func() {
		tests := []struct {
			Value   string
			Success bool
			Dest    int64
		}{
			{"1111", true, 1111},
			{"aaaaa", false, 0},
			{"0", true, 0},
			{"0.0", false, 0},
			{"true", false, 0},
		}

		for _, test := range tests {
			dst, err := Char2Int64(test.Value)
			So(err == nil, ShouldEqual, test.Success)
			if test.Success {
				So(dst, ShouldEqual, test.Dest)
			}
		}
	})
}

func TestMustChar2Int64(t *testing.T) {
	Convey("Cast str must to int32", t, func() {
		tests := []struct {
			Value string
			Dest  int64
		}{
			{"1111", 1111},
			{"0", 0},
		}

		for _, test := range tests {
			dst := MustChar2Int64(test.Value)
			So(dst, ShouldEqual, test.Dest)
		}
	})
}
