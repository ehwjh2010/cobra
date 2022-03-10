package types

import (
	"github.com/ehwjh2010/viper/helper/equal"
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

		So(equal.SliceIntEqual(values, []int{0, 1, 2, 3, 4}), ShouldBeTrue)
		So(set.Size(), ShouldEqual, 5)
	})
}

func TestSimpleSet_Del(t *testing.T) {
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

		So(equal.SliceIntEqual(values, []int{0, 1, 2, 3, 4}), ShouldBeTrue)
		So(set.Size(), ShouldEqual, 5)

		set.Del(0)
		values, _ = set.IntValues()
		So(equal.SliceIntEqual(values, []int{0, 1, 2, 3, 4}), ShouldBeFalse)
		So(equal.SliceIntEqual(values, []int{1, 2, 3, 4}), ShouldBeTrue)
		So(equal.SliceIntEqual(values, []int{0, 1, 2, 3}), ShouldBeFalse)
		So(set.Size(), ShouldEqual, 4)
	})
}

func TestSimpleSet_Update(t *testing.T) {
	Convey("Set add value", t, func() {
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
		So(equal.SliceIntEqual(values, []int{0, 1, 2, 3, 4, 9, 8}), ShouldBeTrue)
	})
}

func TestSimpleSet_Union(t *testing.T) {
	Convey("Set add value", t, func() {
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
		So(equal.SliceIntEqual(values, []int{0, 1, 2, 3, 4, 7, 8, 9, 22}), ShouldBeTrue)
	})
}

func TestSimpleSet_Diff(t *testing.T) {
	Convey("Set add value", t, func() {
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
		So(equal.SliceIntEqual(values, []int{2, 3, 4}), ShouldBeTrue)

		s2 := set2.Diff(set1)
		values2, _ := s2.IntValues()
		So(equal.SliceIntEqual(values2, []int{7, 8, 9, 22}), ShouldBeTrue)
	})
}

func TestSimpleSet_Common(t *testing.T) {
	Convey("Set add value", t, func() {
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
		So(equal.SliceIntEqual(values, []int{0, 1}), ShouldBeTrue)
	})
}

func TestNilSimpleSet(t *testing.T) {
	Convey("Set add value", t, func() {
		var s *SimpleSet

		So(s.IsEmpty(), ShouldBeTrue)
		So(s.Copy(), ShouldBeNil)

	})
}
