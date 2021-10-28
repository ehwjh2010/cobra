package utils

import (
	"io/ioutil"
	"os"
)

//OpenFile 文件不存在则创建, 文件存在, 如果ifExistTrunc为true则清空内容, 否则文件以追加模式进行写操作
//返回的文件指针模式为可读可写
//@param filename 文件路径
//@param ifExistTrunc 文件存在时是否清空内容
func OpenFile(filename string, ifExistTrunc bool) (*os.File, error) {
	flag := os.O_RDWR | os.O_CREATE
	if ifExistTrunc {
		flag |= os.O_TRUNC
	} else {
		flag |= os.O_APPEND
	}

	file, err := os.OpenFile(filename, flag, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

//OpenFileWithAppend 文件不存在则创建, 文件存在, 文件以追加模式进行写操作
//返回的文件指针模式为可读可写
//@param filename 文件路径
func OpenFileWithAppend(filename string) (*os.File, error) {
	return OpenFile(filename, false)
}

//OpenFileWithTrunc 文件不存在则创建, 文件存在, 清空内容
//返回的文件指针模式为可读可写
//@param filename 文件路径
func OpenFileWithTrunc(filename string) (*os.File, error) {
	return OpenFile(filename, true)
}

//ReadFile 只适合读取小文件，不适合读取大文件
//会主动关闭文件对象
func ReadFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	} else {
		return content, nil
	}
}

//WriteFile 写数据到文件, 文件对象会自动关闭
//@param filename 文件路径
//@param data 需要写入的数据
func WriteFile(filename string, data []byte, ifExistTrunc bool) error {
	f, err := OpenFile(filename, ifExistTrunc)
	if err != nil {
		return err
	}

	err = WriteFileWithObj(f, data)
	return err
}

//WriteFileWithObj 写数据到文件, 文件对象会自动关闭
//@param f 文件对象指针
//@param data 需要写入的数据
func WriteFileWithObj(f *os.File, data []byte) error {
	_, err := f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

//WriteFileWithNoClose 写数据到文件, 文件对象不会关闭
//@param f 文件对象指针
//@param data 需要写入的数据
func WriteFileWithNoClose(f *os.File, data []byte) error {
	_, err := f.Write(data)
	return err
}
