package requests

import (
	"errors"
	"github.com/ehwjh2010/viper/global"
	"github.com/levigross/grequests"
	"mime"
	"net/http"
	"strings"
)

var (
	ErrNoDisposition    = errors.New("not found content disposition header")
	ErrNotFoundDstField = errors.New("not found dst field")
	ErrNoCookie         = errors.New("not found cookie")
)

type HTTPResponse struct {
	response   *grequests.Response
	cookies    []*http.Cookie
	cookieFlag bool
}

func NewResponse(response *grequests.Response) *HTTPResponse {
	return &HTTPResponse{response: response}
}

//OK 请求是否OK
func (resp *HTTPResponse) OK() bool {
	return resp.response.Ok
}

//Json 请求体序列化结构体
func (resp *HTTPResponse) Json(dst interface{}) error {
	err := resp.response.JSON(dst)
	return err
}

//Text 返回的文本内容
func (resp *HTTPResponse) Text() string {
	s := resp.response.String()
	return s
}

//Bytes 字节, 适用于小文件, 大文件推荐使用Stream
func (resp *HTTPResponse) Bytes() []byte {
	return resp.response.Bytes()
}

//Status 状态码
func (resp *HTTPResponse) Status() int {
	return resp.response.StatusCode
}

//Headers 返回的Header, 大小写不敏感
func (resp *HTTPResponse) Headers(name string) []string {
	header := resp.response.Header
	return header.Values(name)
}

//Header 返回的Header
func (resp *HTTPResponse) Header(name string) string {
	return resp.response.Header.Get(name)
}

//Cookie 获取Cookie, 没有返回空字符串
func (resp *HTTPResponse) Cookie(name string) (*http.Cookie, error) {
	resp.parseCookie()

	name = strings.ToUpper(name)
	for _, c := range resp.cookies {
		if strings.ToUpper(c.Name) == name {
			return c, nil
		}
	}

	return nil, ErrNoCookie
}

//Cookies 获取Cookies, 没有返回空切片
func (resp *HTTPResponse) Cookies() []*http.Cookie {
	resp.parseCookie()
	return resp.cookies
}

//parseCookie 解析Cookie
func (resp *HTTPResponse) parseCookie() {
	if resp.cookieFlag {
		return
	}

	cookieStr := resp.Header("Cookie")
	if len(cookieStr) <= 0 {
		resp.cookieFlag = true
		return
	}

	cookies := CookieStr2Cookie(cookieStr)
	resp.cookies = cookies
	resp.cookieFlag = true
}

//Error 错误
func (resp *HTTPResponse) Error() error {
	return resp.response.Error
}

//FileName 文件名
func (resp *HTTPResponse) FileName() (string, error) {
	return resp.FileNameWithCustom("filename")
}

//FileNameWithCustom 文件名
func (resp *HTTPResponse) FileNameWithCustom(nameField string) (string, error) {
	contentDisposition := resp.Header(global.ContentDisposition)
	if len(contentDisposition) <= 0 {
		return "", ErrNoDisposition
	}

	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return "", err
	}

	value, exist := params[nameField]
	if !exist {
		return "", ErrNotFoundDstField
	}

	return value, nil
}

//ContentType 获取Content-Type
func (resp *HTTPResponse) ContentType() string {
	return resp.Header(global.ContentType)
}
