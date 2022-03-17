package integer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSliceIntEqual(t *testing.T) {
	Convey("[]int equal", t, func() {
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
			So(SliceIntEqual(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceIntEqualStrict(t *testing.T) {
	Convey("[]int equal", t, func() {
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
			{[]int{1, 3}, []int{3, 1}, false},
		}

		for _, test := range tests {
			So(SliceIntEqualStrict(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceInt32EqualStrict(t *testing.T) {
	Convey("[]int32 equal", t, func() {
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
			{[]int32{1, 3}, []int32{3, 1}, false},
		}

		for _, test := range tests {
			So(SliceInt32EqualStrict(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceInt32Equal(t *testing.T) {
	Convey("[]int32 equal", t, func() {
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
			So(SliceInt32Equal(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceInt64Equal(t *testing.T) {
	Convey("[]int32 equal", t, func() {
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
			So(SliceInt64Equal(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceInt64EqualStrict(t *testing.T) {
	Convey("[]int32 equal", t, func() {
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
			{[]int64{1, 3}, []int64{3, 1}, false},
		}

		for _, test := range tests {
			So(SliceInt64EqualStrict(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}
