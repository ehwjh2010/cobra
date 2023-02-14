package file

import (
	"github.com/ehwjh2010/viper/helper/path"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func TestOpenFile(t *testing.T) {
	Convey("open file", t, func() {
		filename := "test.txt"

		ok, err := path.RemovePath(filename, false)
		if err != nil {
			panic(err)
		}
		So(ok, ShouldBeTrue)

		file, err := OpenFile(filename)
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = file.Close()
		}()

		if _, err := file.WriteString("11\n22\n33\n"); err != nil {
			panic(err)
		}
		exist, err := path.Exists(filename)
		if err != nil {
			panic(err)
		}

		So(exist, ShouldBeTrue)
		ok, err = path.RemovePath(filename, false)
		if err != nil {
			panic(err)
		}

		So(ok, ShouldBeTrue)
	})
}

func TestOpenFileWithTrunc(t *testing.T) {
	Convey("open file with trunc", t, func() {
		filename := "test.txt"

		ok, err := path.RemovePath(filename, false)
		if err != nil {
			panic(err)
		}
		So(ok, ShouldBeTrue)

		cont := "11\n22\n33\n"
		file, err := OpenFile(filename)
		So(err, ShouldBeNil)
		_, err = file.WriteString(cont)
		So(err, ShouldBeNil)
		err = file.Close()
		So(err, ShouldBeNil)

		content, err := ioutil.ReadFile(filename)
		So(err, ShouldBeNil)
		So(string(content), ShouldEqual, cont)

		f, err := OpenFileWithTrunc(filename)
		So(err, ShouldBeNil)
		_ = f.Close()
		bytes, err := ioutil.ReadFile(filename)
		So(err, ShouldBeNil)
		So(bytes, ShouldBeEmpty)
	})
}
