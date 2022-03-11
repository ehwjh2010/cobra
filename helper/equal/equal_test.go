package equal

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

func TestSliceFloat64EqualStrict(t *testing.T) {
	Convey("[]int32 equal", t, func() {
		tests := []struct {
			Value []float64
			Dest  []float64
			OK    bool
		}{
			{nil, []float64{1.11}, false},
			{nil, nil, true},
			{[]float64{1.11}, []float64{1.1, 2.1}, false},
			{[]float64{1.11, 3.11}, []float64{1.11, 2.11}, false},
			{[]float64{1.11, 3.11}, []float64{1.11, 3.11}, true},
			{[]float64{1.11, 3.11}, []float64{3.11, 1.11}, false},
		}

		for _, test := range tests {
			So(SliceFloat64EqualStrict(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}

func TestSliceFloat64Equal(t *testing.T) {
	Convey("[]int32 equal", t, func() {
		tests := []struct {
			Value []float64
			Dest  []float64
			OK    bool
		}{
			{nil, []float64{1.11}, false},
			{nil, nil, true},
			{[]float64{1.11}, []float64{1.1, 2.1}, false},
			{[]float64{1.11, 3.11}, []float64{1.11, 2.11}, false},
			{[]float64{1.11, 3.11}, []float64{1.11, 3.11}, true},
			{[]float64{1.11, 3.11}, []float64{3.11, 1.11}, true},
		}

		for _, test := range tests {
			So(SliceFloat64Equal(test.Value, test.Dest), ShouldEqual, test.OK)
		}
	})
}
