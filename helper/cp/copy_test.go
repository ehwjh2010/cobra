package cp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Prop struct {
	Name string
	Age  int
	Job  interface{}
}

func TestCopyProperties(t *testing.T) {
	Convey("copy slice properties", t, func() {
		a := []int{1, 2, 3, 4, 5}
		b := make([]int, 0)
		CopyProperties(a, &b)
		So(a, ShouldResemble, b)
	})

	Convey("copy array properties", t, func() {
		a := [5]int{1, 2, 3, 4, 5}
		b := [5]int{}
		CopyProperties(a, &b)
		So(a, ShouldResemble, b)
	})

	Convey("copy struct properties", t, func() {
		a := &Prop{
			Name: "Tom",
			Age:  18,
			Job:  "manager",
		}
		b := &Prop{}
		CopyProperties(a, b)
		So(a, ShouldResemble, b)
	})
}
