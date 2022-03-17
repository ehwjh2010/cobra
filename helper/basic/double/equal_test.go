package double

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

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
