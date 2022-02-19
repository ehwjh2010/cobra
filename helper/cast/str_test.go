package cast

import (
	"github.com/ehwjh2010/viper/helper/equal"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStr2Bytes(t *testing.T) {
	Convey("Caset str to bytes", t, func() {
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
			So(equal.SliceBytesEqual(ret, test.Dest), ShouldEqual, test.Success)
		}
	})
}
