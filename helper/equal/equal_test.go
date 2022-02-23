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
			convey.So(SliceByteEqual(test.Value, test.Dest), convey.ShouldEqual, test.OK)
		}
	})
}

func TestSliceStrEqual(t *testing.T) {
	convey.Convey("[]string equal", t, func() {
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
		}

		for _, test := range tests {
			convey.So(SliceStrEqual(test.Value, test.Dest), convey.ShouldEqual, test.OK)
		}
	})
}

func TestSliceIntEqual(t *testing.T) {
	convey.Convey("[]int equal", t, func() {
		tests := []struct {
			Value []int
			Dest  []int
			OK    bool
		}{
			{nil, []int{1}, false},
			{nil, nil, true},
			{[]int{1}, []int{1, 2}, false},
			{[]int{1, 3}, []int{1, 2}, false},
			{[]int{1, 3}, []int{1, 3}, true},
		}

		for _, test := range tests {
			convey.So(SliceIntEqual(test.Value, test.Dest), convey.ShouldEqual, test.OK)
		}
	})
}

func TestSliceInt32Equal(t *testing.T) {
	convey.Convey("[]int32 equal", t, func() {
		tests := []struct {
			Value []int32
			Dest  []int32
			OK    bool
		}{
			{nil, []int32{1}, false},
			{nil, nil, true},
			{[]int32{1}, []int32{1, 2}, false},
			{[]int32{1, 3}, []int32{1, 2}, false},
			{[]int32{1, 3}, []int32{1, 3}, true},
		}

		for _, test := range tests {
			convey.So(SliceInt32Equal(test.Value, test.Dest), convey.ShouldEqual, test.OK)
		}
	})
}

func TestSliceInt64Equal(t *testing.T) {
	convey.Convey("[]int32 equal", t, func() {
		tests := []struct {
			Value []int64
			Dest  []int64
			OK    bool
		}{
			{nil, []int64{1}, false},
			{nil, nil, true},
			{[]int64{1}, []int64{1, 2}, false},
			{[]int64{1, 3}, []int64{1, 2}, false},
			{[]int64{1, 3}, []int64{1, 3}, true},
		}

		for _, test := range tests {
			convey.So(SliceInt64Equal(test.Value, test.Dest), convey.ShouldEqual, test.OK)
		}
	})
}
