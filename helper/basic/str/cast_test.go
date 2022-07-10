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

func TestStrSlice2Any(t *testing.T) {
	Convey("Cast slice string to any", t, func() {
		foo := []string{"a", "b", "c", "d"}
		temp := SliceStr2Any(foo)
		slice2Str, err := SliceAny2Char(temp)
		So(err, ShouldBeNil)
		So(slice2Str, ShouldResemble, foo)

		v := SliceStr2Any(nil)
		So(v, ShouldBeNil)

		tmp := make([]interface{}, 3)
		tmp[0] = 1
		tmp[0] = "1"
		tmp[0] = true

		slice2Str, err = SliceAny2Char(tmp)
		So(err, ShouldNotBeNil)
		So(slice2Str, ShouldBeEmpty)

		str, err := SliceAny2Char(nil)
		So(err, ShouldBeNil)
		So(str, ShouldBeEmpty)
	})
}
