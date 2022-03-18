package file

import "os"

// openFile 文件不存在则创建, 文件存在, 如果ifExistTrunc为true则清空内容, 否则文件以追加模式进行写操作
// 返回的文件指针模式为可读可写
// filename 文件路径
// ifExistTrunc 文件存在时是否清空内容
func openFile(filename string, ifExistTrunc bool) (*os.File, error) {
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

// OpenFile 文件不存在则创建, 文件存在, 文件以追加模式进行写操作
// 返回的文件指针模式为可读可写
// filename 文件路径
func OpenFile(filename string) (*os.File, error) {
	return openFile(filename, false)
}

// OpenFileWithTrunc 文件不存在则创建, 文件存在, 清空内容
// 返回的文件指针模式为可读可写
// filename 文件路径
func OpenFileWithTrunc(filename string) (*os.File, error) {
	return openFile(filename, true)
}
