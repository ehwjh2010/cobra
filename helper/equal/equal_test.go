package equal

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSliceBytesEqual(t *testing.T) {
	convey.Convey("[]byte equal", t, func() {
		tests := []struct {
			Value []byte
			Dest  []byte
			OK    bool
		}{
			{nil, []byte{'1'}, false},
			{nil, nil, true},
			{[]byte{'1'}, []byte{'1', '2'}, false},
			{[]byte{'1', '3'}, []byte{'1', '2'}, false},
			{[]byte{'1', '3'}, []byte{'1', '3'}, true},
		}

		for _, test := range tests {
			convey.So(SliceBytesEqual(test.Value, test.Dest), convey.ShouldEqual, test.OK)
		}
	})
}
