package requests

import (
	"io"
)

type FileUpload struct {
	FileName     string        //文件名
	FileContents io.ReadCloser //文件流
	FieldName    string        //上传文件表单 文件字段
	FileMime     string        //文件类型, 默认是application/octet-stream
}
