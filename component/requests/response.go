package requests

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"mime"
	"net/http"

	"github.com/ehwjh2010/viper/constant"
	"github.com/ehwjh2010/viper/helper/cookies"
)

var (
	ErrNoDisposition    = errors.New("not found content disposition header")
	ErrNotFoundDstField = errors.New("not found dst field")
)

type HTTPResponse struct {
	cookies  []*http.Cookie
	response *resty.Response
}

func NewResponse(response *resty.Response) *HTTPResponse {
	return &HTTPResponse{response: response}
}

// OK 状态码是否在200~300之间.
func (resp *HTTPResponse) OK() bool {
	return resp.response.IsSuccess()
}

// Json 请求体序列化结构体.
func (resp *HTTPResponse) Json(dst interface{}) error {
	err := json.Unmarshal(resp.response.Body(), dst)
	return err
}

// Text 返回的文本内容.
func (resp *HTTPResponse) Text() string {
	s := resp.response.String()
	return s
}

// StatusCode 状态码.
func (resp *HTTPResponse) StatusCode() int {
	return resp.response.StatusCode()
}

// Header 返回的Header.
func (resp *HTTPResponse) Header(name string) string {
	value := resp.response.Header()[name]
	if len(value) == 0 {
		return ""
	}

	return value[0]
}

// Cookie 获取Cookie, 没有返回空字符串.
func (resp *HTTPResponse) Cookie(name string) (*http.Cookie, error) {
	cookieParser := cookies.NewCookieParser()
	cookie, err := cookieParser.GetDestFromCookies(resp.response.Cookies(), name)
	return cookie, err
}

// Cookies 获取Cookies, 没有返回空切片.
func (resp *HTTPResponse) Cookies() []*http.Cookie {
	return resp.response.Cookies()
}

// FileName 文件名.
func (resp *HTTPResponse) FileName() (string, error) {
	return resp.FileNameWithCustom("filename")
}

// FileNameWithCustom 文件名.
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

// ContentType 获取Content-Type.
func (resp *HTTPResponse) ContentType() string {
	return resp.Header(constant.ContentType)
}
