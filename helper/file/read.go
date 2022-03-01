package file

import (
	"io/ioutil"
)

// ReadFile 只适合读取小文件，不适合读取大文件
// 会主动关闭文件对象
func ReadFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	} else {
		return content, nil
	}
}
