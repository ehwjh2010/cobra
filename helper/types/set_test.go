package types

import (
	"fmt"
	"github.com/ehwjh2010/viper/helper/basic/double"
	"github.com/ehwjh2010/viper/helper/basic/integer"
	"github.com/ehwjh2010/viper/helper/basic/str"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSimpleSet_Add(t *testing.T) {
	Convey("Set add value", t, func() {
		set := NewSimpleSet()
		set.Add(0)
		set.Add(0)
		set.Add(1)
		set.Add(2)
		set.Add(3)
		set.Add(4)
		set.Add(4)

		values, _ := set.IntValues()

		So(integer.SliceIntEqual(values, []int{0, 1, 2, 3, 4}), ShouldBeTrue)
		So(set.Size(), ShouldEqual, 5)

		var set2 *Set
		set2.Add(1)
		set2.Add(2)
		set2.Add(3)
		So(set2.IsEmpty(), ShouldBeTrue)
		set2.Del("art")
		So(set2.IsEmpty(), ShouldBeTrue)

		set3 := NewSimpleSet()
		set3.Add("1")
		set3.Add(2)
		set3.Add(true)
		intValues, err := set3.IntValues()
		So(err, ShouldNotBeNil)
		So(intValues, ShouldBeEmpty)
	})
}

func TestSimpleSet_Int64Values(t *testing.T) {
	Convey("Set set int64 values", t, func() {
		foo := []int64{1, 2, 3, 4, 5, 1}
		set := NewSimpleSet()
		set.Update(integer.SliceInt64TAny(foo)...)
		values, err := set.Int64Values()
		So(err, ShouldBeNil)
		So(integer.SliceInt64Equal(values, []int64{1, 2, 3, 4, 5}), ShouldBeTrue)

		set.Add(true)
		set.Add("art")
		values, err = set.Int64Values()
		So(err, ShouldNotBeNil)
		So(values, ShouldBeEmpty)

	})
}

func TestSimpleSet_Int32Values(t *testing.T) {
	Convey("Set set int32 values", t, func() {
		var vs = []int32{0, 1, 2, 3, 4}
		set := NewSimpleSet()
		vsi := integer.SliceInt32TAny(vs)
		set.Update(vsi...)

		values, _ := set.Int32Values()
		fmt.Println(values)

		So(integer.SliceInt32Equal(values, []int32{0, 1, 2, 3, 4}), ShouldBeTrue)
		So(set.Size(), ShouldEqual, 5)

		set2 := NewSimpleSet()
		set2.Add(1)
		set2.Add("22")
		int32Values, err := set2.Int32Values()
		So(int32Values, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func TestSimpleSet_StrValues(t *testing.T) {
	Convey("Set set string values", t, func() {
		set := NewSimpleSet()
		set.Add("0")
		set.Add("0")
		set.Add("1")
		set.Add("2")
		set.Add("3")
		set.Add("4")
		set.Add("4")

		values, _ := set.StrValues()

		So(str.SliceStrEqual(values, []string{"0", "1", "2", "3", "4"}), ShouldBeTrue)
		So(set.Size(), ShouldEqual, 5)

		set2 := NewSimpleSet()
		set2.Add(1)
		set2.Add("22")
		strValues, err := set2.StrValues()
		So(strValues, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func TestSimpleSet_Float64Values(t *testing.T) {
	Convey("Set set float46 values", t, func() {
		set := NewSimpleSet()
		set.Add(0.0)
		set.Add(1.1)
		set.Add(2.2)
		set.Add(3.3)
		set.Add(4.4)
		set.Add(5.55)
		set.Add(4.0)

		values, _ := set.Float64Values()

		So(double.SliceFloat64Equal(values, []float64{0, 1.1, 2.2, 3.3, 4.4, 4.0, 5.55}), ShouldBeTrue)
		So(set.Size(), ShouldEqual, 7)

		set2 := NewSimpleSet()
		set2.Add(1.22)
		set2.Add("22")
		float64Values, err := set2.Float64Values()
		So(float64Values, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func TestSimpleSet_Del(t *testing.T) {
	Convey("Set set del", t, func() {
		set := NewSimpleSet()
		set.Add(0)
		set.Add(0)
		set.Add(1)
		set.Add(2)
		set.Add(3)
		set.Add(4)
		set.Add(4)

		values, _ := set.IntValues()

		So(integer.SliceIntEqual(values, []int{0, 1, 2, 3, 4}), ShouldBeTrue)
		So(set.Size(), ShouldEqual, 5)

		set.Del(0)
		values, _ = set.IntValues()
		So(integer.SliceIntEqual(values, []int{0, 1, 2, 3, 4}), ShouldBeFalse)
		So(integer.SliceIntEqual(values, []int{1, 2, 3, 4}), ShouldBeTrue)
		So(integer.SliceIntEqual(values, []int{0, 1, 2, 3}), ShouldBeFalse)
		So(set.Size(), ShouldEqual, 4)
	})
}

func TestSimpleSet_Update(t *testing.T) {
	Convey("Set set update", t, func() {
		set := NewSimpleSet()
		set.Add(0)
		set.Add(1)
		set.Add(2)
		set.Add(3)
		set.Add(4)

		var p = []interface{}{9, 0, 8, 1}
		set.Update(p...)

		values, err := set.IntValues()
		So(err, ShouldBeNil)
		So(integer.SliceIntEqual(values, []int{0, 1, 2, 3, 4, 9, 8}), ShouldBeTrue)

		var set2 *Set
		set2.Update(1, 2, 3, "A")
		So(set2.IsEmpty(), ShouldBeTrue)
	})
}

func TestSimpleSet_Union(t *testing.T) {
	Convey("Set set union", t, func() {
		set1 := NewSimpleSet()
		set1.Add(0)
		set1.Add(1)
		set1.Add(2)
		set1.Add(3)
		set1.Add(4)

		set2 := NewSimpleSet()
		set2.Add(8)
		set2.Add(9)
		set2.Add(7)
		set2.Add(22)
		set2.Add(0)
		set2.Add(1)

		s := set1.Union(set2)
		values, _ := s.IntValues()
		So(integer.SliceIntEqual(values, []int{0, 1, 2, 3, 4, 7, 8, 9, 22}), ShouldBeTrue)
	})
}

func TestSimpleSet_Diff(t *testing.T) {
	Convey("Set set diff", t, func() {
		set1 := NewSimpleSet()
		set1.Add(0)
		set1.Add(1)
		set1.Add(2)
		set1.Add(3)
		set1.Add(4)

		set2 := NewSimpleSet()
		set2.Add(8)
		set2.Add(9)
		set2.Add(7)
		set2.Add(22)
		set2.Add(0)
		set2.Add(1)

		s := set1.Diff(set2)
		values, _ := s.IntValues()
		So(integer.SliceIntEqual(values, []int{2, 3, 4}), ShouldBeTrue)

		s2 := set2.Diff(set1)
		values2, _ := s2.IntValues()
		So(integer.SliceIntEqual(values2, []int{7, 8, 9, 22}), ShouldBeTrue)

		set3 := NewSimpleSet()
		So(set3.Diff(set1).IsEmpty(), ShouldBeTrue)
		So(set1.Diff(set3).IsNotEmpty(), ShouldBeTrue)
	})
}

func TestSimpleSet_Common(t *testing.T) {
	Convey("Set set common", t, func() {
		set1 := NewSimpleSet()
		set1.Add(0)
		set1.Add(1)
		set1.Add(2)
		set1.Add(3)
		set1.Add(4)

		set2 := NewSimpleSet()
		set2.Add(8)
		set2.Add(9)
		set2.Add(7)
		set2.Add(22)
		set2.Add(0)
		set2.Add(1)

		s := set1.Common(set2)
		values, _ := s.IntValues()
		So(integer.SliceIntEqual(values, []int{0, 1}), ShouldBeTrue)

		set3 := NewSimpleSet()
		So(set1.Common(set3).IsEmpty(), ShouldBeTrue)
		So(set3.Common(set1).IsEmpty(), ShouldBeTrue)
	})
}

func TestNilSimpleSet(t *testing.T) {
	Convey("Set nil set", t, func() {
		var s *Set
		var s2 *Set

		So(s.Values(), ShouldBeEmpty)
		So(s.IsEmpty(), ShouldBeTrue)
		So(s.Copy().IsEmpty(), ShouldBeTrue)
		diff := s.Diff(s2)
		So(diff.IsEmpty(), ShouldBeTrue)
		So(diff.IsNotEmpty(), ShouldBeFalse)

		So(s.Size(), ShouldEqual, 0)
		s.Add(1)
		So(s.Has(1), ShouldBeFalse)

		values, err := s.IntValues()
		So(err, ShouldBeNil)
		So(values, ShouldBeEmpty)

		int32Values, err := s.Int32Values()
		So(err, ShouldBeNil)
		So(int32Values, ShouldBeEmpty)

		int64Values, err := s.Int64Values()
		So(err, ShouldBeNil)
		So(int64Values, ShouldBeEmpty)

		strValues, err := s.StrValues()
		So(err, ShouldBeNil)
		So(strValues, ShouldBeEmpty)

		float64Values, err := s.Float64Values()
		So(err, ShouldBeNil)
		So(float64Values, ShouldBeEmpty)
	})
}
