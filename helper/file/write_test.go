package file

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWriteFile(t *testing.T) {
	Convey("Write file", t, func() {
		err := WriteFile("write.txt", []byte("test1\ntest2\n"), true)
		if err != nil {
			panic(err)
		}
	})
}

func TestWriteFileWithoutClose(t *testing.T) {
	Convey("Write file without close", t, func() {
		f, err := OpenFileWithTrunc("write_no_close.txt")
		if err != nil {
			panic(err)
		}

		defer f.Close()
		err = WriteFileWithoutClose(f, []byte("test1\ntest2\n"))
		if err != nil {
			panic(err)
		}
	})
}
