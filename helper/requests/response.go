package requests

import (
	"github.com/ehwjh2010/viper/helper/object"
	"github.com/levigross/grequests"
	"io"
)

type Response struct {
	*grequests.Response
}

//OK 请求是否OK
func (resp Response) OK() bool {
	return resp.Response.Ok
}

//Json 请求体序列化结构体
func (resp Response) Json(dst interface{}) error {
	err := resp.Response.JSON(dst)
	return err
}

//Text 返回的文本内容
func (resp Response) Text() string {
	s := resp.Response.String()
	return s
}

//Bytes 字节, 适用于小文件, 大文件推荐使用Stream
func (resp *Response) Bytes() []byte {
	return resp.Response.Bytes()
}

//Status 状态码
func (resp Response) Status() int {
	return resp.Response.StatusCode
}

//Headers 返回的Header, 大小写不敏感
func (resp Response) Headers(name string) []string {
	header := resp.Response.Header
	return header.Values(name)
}

//Header 返回的Header
func (resp Response) Header(name string) string {
	return resp.Response.Header.Get(name)
}

//Cookie 获取Cookie
func (resp Response) Cookie(name string) string {
	//TODO 从header解析Cookie
	resp.Header("Cookie")

	return ""
}

//Error 错误
func (resp Response) Error() error {
	return resp.Response.Error
}

//Stream 文件流, 会利用磁盘处理大文件, 避免OOM
func (resp Response) Stream() io.ReadCloser {
	//TODO 处理大文件流
	return nil
}

//FileName 文件名
func (resp Response) FileName() string {
	return resp.FileNameWithCustom("filename")
}

//FileNameWithCustom 文件名
func (resp Response) FileNameWithCustom(nameField string) string {
	//TODO 解析header的文件名
	return ""
}

//toReqOptions 转换RequestOptions
func (r Request) toReqOptions() *grequests.RequestOptions {
	rOpt := &grequests.RequestOptions{
		Headers:        r.Header,
		Params:         r.Params,
		Data:           r.Form,
		JSON:           r.Json,
		Cookies:        r.Cookie,
		UserAgent:      r.UserAgent,
		RequestTimeout: r.Timeout,
	}

	if len(r.Files) > 0 {
		files := make([]grequests.FileUpload, len(r.Files))
		object.CopyProperties(r.Files, files)
		rOpt.Files = files
	}

	return rOpt
}
