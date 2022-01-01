package requests

import "io"

type FileUpload struct {
	FileName  string        //文件名
	File      io.ReadCloser //文件流
	FieldName string        //上传文件表单 文件字段
	FileMime  string        //文件类型, 默认是application/octet-stream
}

//handleFileMime 设置FileMime
func (fileUpload *FileUpload) handleFileMime() {
	//TODO 自动根据文件名处理Mime, 默认使用application/octet-stream
}
