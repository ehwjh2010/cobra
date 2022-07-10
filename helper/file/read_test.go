package file

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	Convey("Read file", t, func() {
		_, err := ReadFile("test2.txt")
		So(errors.Is(err, os.ErrNotExist), ShouldBeTrue)

		content, err := ReadFile("read.txt")
		if err != nil {
			t.Error(err)
		} else {
			So(string(content), ShouldEqual, "aaa333\nbbb444")
		}
	})
}
