package requests

import (
	"io"

	"github.com/ehwjh2010/viper/helper/cp"
	"github.com/levigross/grequests"
)

type FileUpload struct {
	FileName     string        //文件名
	FileContents io.ReadCloser //文件流
	FieldName    string        //上传文件表单 文件字段
	FileMime     string        //文件类型, 默认是application/octet-stream
}

// toInternal 转化为grequests.FileUpload.
func (fileUpload *FileUpload) toInternal() grequests.FileUpload {

	var gUpload grequests.FileUpload

	cp.CopyProperties(fileUpload, &gUpload)

	return gUpload
}

// BatchFileUpload2Internal 批量转化.
func BatchFileUpload2Internal(files []*FileUpload) []grequests.FileUpload {
	if files == nil {
		return nil
	}

	result := make([]grequests.FileUpload, len(files))
	for index, file := range files {
		if file == nil {
			continue
		}
		internal := file.toInternal()
		result[index] = internal
	}

	return result
}
