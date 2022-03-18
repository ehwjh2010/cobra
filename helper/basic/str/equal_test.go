package str

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSliceBytesEqual(t *testing.T) {
	Convey("[]byte equal", t, func() {
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
			So(SliceByteEqual(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceStrEqual(t *testing.T) {
	Convey("[]string equal", t, func() {
		tests := []struct {
			Value []string
			Dest  []string
			OK    bool
		}{
			{nil, []string{"1"}, false},
			{nil, nil, true},
			{[]string{"1"}, []string{"1", "2"}, false},
			{[]string{"1", "3"}, []string{"1", "2"}, false},
			{[]string{"1", "3"}, []string{"1", "3"}, true},
			{[]string{"1", "3"}, []string{"3", "1"}, true},
		}

		for _, test := range tests {
			So(SliceStrEqual(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceStrEqualStrict(t *testing.T) {
	Convey("[]string equal", t, func() {
		tests := []struct {
			Value []string
			Dest  []string
			OK    bool
		}{
			{nil, []string{"1"}, false},
			{nil, nil, true},
			{[]string{"1"}, []string{"1", "2"}, false},
			{[]string{"1", "3"}, []string{"1", "2"}, false},
			{[]string{"1", "3"}, []string{"1", "3"}, true},
			{[]string{"1", "3"}, []string{"3", "1"}, false},
		}

		for _, test := range tests {
			So(SliceStrEqualStrict(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}
