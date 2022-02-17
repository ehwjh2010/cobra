package cast

import (
	"github.com/ehwjh2010/viper/helper/equal"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	Convey("Cast str to byte", t, func() {
		tests := []struct {
			Value string
			Dest  []byte
		}{
			{"aaa", []byte("aaa")},
			{"bool", []byte("bool")},
			{"", []byte("")},
		}

		var ret []byte
		for _, test := range tests {
			ret = Str2Bytes(test.Value)
			r := equal.SliceBytesEqual(ret, test.Dest)
			So(r, ShouldBeTrue)
		}
	})
}
