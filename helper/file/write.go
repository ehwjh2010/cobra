package file

import "os"

// WriteFile 写数据到文件, 文件对象会自动关闭
// filename 文件路径
// data 需要写入的数据
func WriteFile(filename string, data []byte, ifExistTrunc bool) error {
	f, err := openFile(filename, ifExistTrunc)
	if err != nil {
		return err
	}

	err = WriteFileWithObj(f, data)
	return err
}

// WriteFileWithObj 写数据到文件, 文件对象会自动关闭
// f 文件对象指针
// data 需要写入的数据
func WriteFileWithObj(f *os.File, data []byte) error {
	_, err := f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

// WriteFileWithNoClose 写数据到文件, 文件对象不会关闭
// f 文件对象指针
// data 需要写入的数据
func WriteFileWithNoClose(f *os.File, data []byte) error {
	_, err := f.Write(data)
	return err
}
