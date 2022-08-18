package types

import "net/http"

type Response struct {
	Cookie     []*http.Cookie
	Header     map[string]string
	Result     interface{}
	StatusCode int
}

type RespOpt func(response *Response)

func NewResp(result Result, args ...RespOpt) *Response {
	response := &Response{
		Result:     result,
		StatusCode: http.StatusOK,
	}

	for _, arg := range args {
		arg(response)
	}

	return response
}

// NewCustomResp 自定义返回, 配合 CustomResponse 使用.
func NewCustomResp(statusCode int, code int, message string, data interface{}) *Response {
	result := NewResult(data, ResultWithCode(code), ResultWithMessage(message))
	response := NewResp(result, RespWithStatus(statusCode))
	return response
}

// SetStatusCode 设置状态码.
func (r *Response) SetStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

// AddHeader 添加请求头, 如果值为"", 则被视为删除该头部.
func (r *Response) AddHeader(key, value string) *Response {
	r.Header[key] = value
	return r
}

// AddHeaders 添加请求头.
func (r *Response) AddHeaders(headers map[string]string) *Response {
	for k, v := range headers {
		r.Header[k] = v
	}

	return r
}

// AddCookie 添加Cookie.
func (r *Response) AddCookie(cookies ...*http.Cookie) *Response {
	for _, cookie := range cookies {
		if cookie != nil {
			r.Cookie = append(r.Cookie, cookie)
		}
	}

	return r
}

// RespWithHeader 添加header.
func RespWithHeader(key, value string) RespOpt {
	return func(response *Response) {
		response.AddHeader(key, value)
	}
}

// RespWithHeaders 添加多个header.
func RespWithHeaders(header map[string]string) RespOpt {
	return func(response *Response) {
		response.AddHeaders(header)
	}
}

// RespWithCookies 设置Cookie.
func RespWithCookies(cookies ...*http.Cookie) RespOpt {
	return func(response *Response) {
		response.AddCookie(cookies...)
	}
}

// RespWithStatus 设置状态码.
func RespWithStatus(statusCode int) RespOpt {
	return func(response *Response) {
		response.StatusCode = statusCode
	}
}
