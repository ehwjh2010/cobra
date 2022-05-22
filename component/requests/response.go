package requests

import (
	"errors"
	"github.com/ehwjh2010/viper/constant"
	"mime"
	"net/http"
	"strings"

	"github.com/levigross/grequests"
)

var (
	ErrNoDisposition    = errors.New("not found content disposition header")
	ErrNotFoundDstField = errors.New("not found dst field")
	ErrNoCookie         = errors.New("not found cookie")
)

type HTTPResponse struct {
	response        *grequests.Response
	cookies         []*http.Cookie
	parseCookieFlag bool
}

func NewResponse(response *grequests.Response) *HTTPResponse {
	return &HTTPResponse{response: response}
}

// OK 状态码是否在200~300之间
func (resp *HTTPResponse) OK() bool {
	return resp.response.Ok
}

// Json 请求体序列化结构体
func (resp *HTTPResponse) Json(dst interface{}) error {
	err := resp.response.JSON(dst)
	return err
}

// Text 返回的文本内容
func (resp *HTTPResponse) Text() string {
	s := resp.response.String()
	return s
}

// Bytes 字节, 适用于小文件, 大文件推荐使用Stream
func (resp *HTTPResponse) Bytes() []byte {
	return resp.response.Bytes()
}

// Status 状态码
func (resp *HTTPResponse) Status() int {
	return resp.response.StatusCode
}

// Header 返回的Header
func (resp *HTTPResponse) Header(name string) string {
	return resp.response.Header.Get(name)
}

// Cookie 获取Cookie, 没有返回空字符串
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

// Cookies 获取Cookies, 没有返回空切片
func (resp *HTTPResponse) Cookies() []*http.Cookie {
	resp.parseCookie()
	return resp.cookies
}

// parseCookie 解析Cookie
func (resp *HTTPResponse) parseCookie() {
	if resp.parseCookieFlag {
		return
	}

	cookieStr := resp.Header("Cookie")
	if len(cookieStr) <= 0 {
		resp.parseCookieFlag = true
		return
	}

	cookies := CookieStr2Cookie(cookieStr)
	resp.cookies = cookies
	resp.parseCookieFlag = true
}

// FileName 文件名
func (resp *HTTPResponse) FileName() (string, error) {
	return resp.FileNameWithCustom("filename")
}

// FileNameWithCustom 文件名
func (resp *HTTPResponse) FileNameWithCustom(nameField string) (string, error) {
	contentDisposition := resp.Header(constant.ContentDisposition)
	if len(contentDisposition) <= 0 {
		return "", ErrNoDisposition
	}

	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return "", err
	}

	value, exists := params[nameField]
	if !exists {
		return "", ErrNotFoundDstField
	}

	return value, nil
}

// ContentType 获取Content-Type
func (resp *HTTPResponse) ContentType() string {
	return resp.Header(constant.ContentType)
}
