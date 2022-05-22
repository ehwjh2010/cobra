package path

import (
	"errors"
	"github.com/ehwjh2010/viper/constant"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var ErrPathAlreadyExist = errors.New("path already exist")
var ErrPathNoExist = errors.New("path no exist")
var ErrInvalidPath = errors.New("invalid path")

// EnsurePathExist 确认文件或文件夹是否存在
func EnsurePathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}

// MakeDir 创建单一目录, 不支持创建多级目录
// path 路径
// exist_no_error 路径已存在时是否返回错误
func MakeDir(path string, existReturnError bool) (err error) {
	if path == "" {
		return ErrInvalidPath
	}

	exist, err := EnsurePathExist(path)
	if err != nil {
		return
	}

	if exist {
		if existReturnError {
			return ErrPathAlreadyExist
		} else {
			return
		}
	} else {
		err = os.Mkdir(path, 0777)
	}
	return
}

// MakeDirIfNotPresent 目录不存在, 则创建; 存在则不操作
// path 路径
func MakeDirIfNotPresent(path string) error {
	return MakeDir(path, false)
}

// RemovePath 完全删除文件夹或文件, 对于文件夹包括子文件以及子文件夹
// path 路径
// noExistReturnError 路径不存在时是否返回错误
func RemovePath(path string, noExistReturnError bool) (bool, error) {
	if path == "" {
		return false, ErrInvalidPath
	}

	exists, err := EnsurePathExist(path)
	if err != nil {
		return false, err
	}

	if !exists {
		if noExistReturnError {
			return false, ErrPathNoExist
		} else {
			return true, nil
		}

	} else {
		err := os.RemoveAll(path)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
}

// JoinPath 路径拼接
func JoinPath(paths ...string) string {
	p := filepath.Join(paths...)
	return p
}

//Relative2Abs 相对路径转化绝对路径
func Relative2Abs(relativePath string) (string, error) {

	if relativePath == "" {
		return "", nil
	}

	if strings.HasPrefix(relativePath, constant.HomeShortCut) {
		home := os.Getenv("HOME")
		relativePath = strings.Replace(relativePath, constant.HomeShortCut, home, 1)
	}

	absPath, err := filepath.Abs(relativePath)

	if err != nil {
		return "", err
	} else {
		return absPath, nil
	}
}

// PathSplit 路径分割, 返回目录以及文件名
func PathSplit(path string) (string, string) {
	dirName := filepath.Dir(path)
	fileName := filepath.Base(path)
	return dirName, fileName
}

// MakeDirs 创建多级目录
func MakeDirs(path ...string) error {
	tmp := JoinPath(path...)

	if tmp == "" {
		return nil
	}

	return os.MkdirAll(tmp, os.ModePerm)
}
