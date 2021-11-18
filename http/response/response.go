package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Cookies    []*http.Cookie
	Header     map[string]string
	Result     *Result
	StatusCode int
}

type RespOpt func(response *Response)

func NewResp(result *Result, args ...RespOpt) *Response {
	response := &Response{
		Result:     result,
		StatusCode: http.StatusOK,
	}

	for _, arg := range args {
		arg(response)
	}

	return response
}

//NewCustomResp 自定义返回, 配合 CustomResponse 使用
func NewCustomResp(statusCode int, code int, message string, data interface{}) *Response {
	result := NewResult(data, ResultWithCode(code), ResultWithMessage(message))
	response := NewResp(result, RespWithStatusCode(statusCode))
	return response
}

//SetStatusCode 设置状态码
func (r *Response) SetStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

//AddHeader 添加请求头, 如果值为"", 则被视为删除该头部
func (r *Response) AddHeader(key, value string) *Response {
	r.Header[key] = value
	return r
}

//AddHeaders 添加请求头
func (r *Response) AddHeaders(headers map[string]string) *Response {
	for k, v := range headers {
		r.Header[k] = v
	}

	return r
}

//AddCookie 添加Cookie
func (r *Response) AddCookie(cookie ...*http.Cookie) *Response {
	if cookie != nil {
		r.Cookies = append(r.Cookies, cookie...)
	}

	return r
}

//RespWithHeader 添加header
func RespWithHeader(key, value string) RespOpt {
	return func(response *Response) {
		response.AddHeader(key, value)
	}
}

//RespWithHeaders 添加多个header
func RespWithHeaders(header map[string]string) RespOpt {
	return func(response *Response) {
		response.AddHeaders(header)
	}
}

//RespWithCookies 设置Cookie
func RespWithCookies(cookies ...*http.Cookie) RespOpt {
	return func(response *Response) {
		response.AddCookie(cookies...)
	}
}

//RespWithStatusCode 设置状态码
func RespWithStatusCode(statusCode int) RespOpt {
	return func(response *Response) {
		response.StatusCode = statusCode
	}
}

//Success 业务正常运行, 状态码200
func Success(c *gin.Context, data interface{}, args ...RespOpt) {
	result := NewResult(data)

	response := NewResp(result, args...)

	baseResponse(c, response)
}

//InvalidRequest 无效入参, 状态码400
func InvalidRequest(c *gin.Context, message string) {
	result := NewResult(nil, ResultWithMessage(message), ResultWithCode(InvalidParams))

	response := NewResp(result, RespWithStatusCode(http.StatusBadRequest))

	baseResponse(c, response)
}

//InvalidRequestWithData 与InvalidRequest功能一致, 但是可以返回数据
func InvalidRequestWithData(c *gin.Context, message string, data interface{}) {
	result := NewResult(data, ResultWithMessage(message), ResultWithCode(InvalidParams))

	response := NewResp(result, RespWithStatusCode(http.StatusBadRequest))

	baseResponse(c, response)
}

//Fail 请求正常, 状态码200, 但是业务流不正常
func Fail(c *gin.Context, code int, message string, args ...RespOpt) {
	result := NewResult(nil, ResultWithCode(code), ResultWithMessage(message))

	response := NewResp(result, args...)

	baseResponse(c, response)
}

//FailWithData 功能与Fail一致, 只不过FailWithData可以返回数据
func FailWithData(c *gin.Context, code int, message string, data interface{}, args ...RespOpt) {
	result := NewResult(data, ResultWithCode(code), ResultWithMessage(message))

	response := NewResp(result, args...)

	baseResponse(c, response)
}

func CustomResponse(c *gin.Context, response *Response) {
	baseResponse(c, response)
}

func baseResponse(c *gin.Context, response *Response) {
	if len(response.Header) > 0 {
		for key, value := range response.Header {
			c.Header(key, value)
		}
	}

	if response.Cookies != nil {
		for _, cookie := range response.Cookies {
			if cookie == nil {
				continue
			}
			c.SetCookie(
				cookie.Name,
				cookie.Value,
				cookie.MaxAge,
				cookie.Path,
				cookie.Domain,
				cookie.Secure,
				cookie.HttpOnly)
		}
	}

	c.JSON(response.StatusCode, response.Result)
}
