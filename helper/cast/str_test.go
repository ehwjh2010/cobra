package cast

import (
	"github.com/ehwjh2010/viper/helper/equal"
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
			ret = Str2Bytes(test.Value)
			So(equal.SliceByteEqual(ret, test.Dest), ShouldEqual, test.Success)
		}
	})
}

func TestAny2String(t *testing.T) {
	Convey("Cast any to string", t, func() {
		var foo interface{}
		origin := "1"
		foo = origin

		any2String := MustAny2String(foo)
		So(any2String, ShouldEqual, origin)

		foo = 2
		any2String, err := Any2String(foo)
		So(err, ShouldNotBeNil)
		So(any2String, ShouldBeEmpty)
	})
}

func TestStr2Any(t *testing.T) {
	Convey("Cast string to any", t, func() {
		foo := "art"
		So(MustAny2String(Str2Any(foo)), ShouldEqual, foo)
	})
}

func TestStrSlice2Any(t *testing.T) {
	Convey("Cast slice string to any", t, func() {
		foo := []string{"a", "b", "c", "d"}
		temp := StrSlice2Any(foo)
		slice2Str, err := AnySlice2Str(temp)
		So(err, ShouldBeNil)
		So(slice2Str, ShouldResemble, foo)

		v := StrSlice2Any(nil)
		So(v, ShouldBeNil)

		tmp := make([]interface{}, 3)
		tmp[0] = 1
		tmp[0] = "1"
		tmp[0] = true

		slice2Str, err = AnySlice2Str(tmp)
		So(err, ShouldNotBeNil)
		So(slice2Str, ShouldBeEmpty)

		str, err := AnySlice2Str(nil)
		So(err, ShouldBeNil)
		So(str, ShouldBeEmpty)
	})
}
