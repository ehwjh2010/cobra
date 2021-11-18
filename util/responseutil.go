package util

import "github.com/gin-gonic/gin"
import "net/http"

const (
	OK      = 0
	SUCCESS = "Success"

	InvalidParams = 10
)

type (
	Result struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	Response struct {
		Cookies    []*http.Cookie
		Header     map[string]string
		Result     *Result
		StatusCode int
	}
)

//-------------------Result-----------------------

func NewResult(data interface{}, args ...ResultOpt) *Result {
	result := &Result{
		Code:    OK,
		Message: SUCCESS,
		Data:    data,
	}
	for _, arg := range args {
		arg(result)
	}

	return result
}

type ResultOpt func(response *Result)

func ResultWithCode(code int) ResultOpt {
	return func(response *Result) {
		response.Code = code
	}
}

func ResultWithMessage(msg string) ResultOpt {
	return func(response *Result) {
		response.Message = msg
	}
}

func ResultWithData(data interface{}) ResultOpt {
	return func(response *Result) {
		response.Data = data
	}
}

//-------------------Response-----------------------

type ResponseOpt func(response *Response)

func NewResponse(result *Result, args ...ResponseOpt) *Response {
	response := &Response{
		Result:     result,
		StatusCode: http.StatusOK,
	}

	for _, arg := range args {
		arg(response)
	}

	return response
}

//SetStatusCode 设置状态码
func (r *Response) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

//AddHeader 添加请求头
func (r *Response) AddHeader(key, value string) {
	r.Header[key] = value
}

//AddHeaders 添加请求头
func (r *Response) AddHeaders(headers map[string]string) {
	for k, v := range headers {
		r.Header[k] = v
	}
}

//AddCookie 添加Cookie
func (r *Response) AddCookie(cookie ...*http.Cookie) {
	if cookie == nil {
		return
	}

	r.Cookies = append(r.Cookies, cookie...)
}

func ResponseWithHeader(header map[string]string) ResponseOpt {
	return func(response *Response) {
		response.Header = header
	}
}

func ResponseWithCookie(cookie []*http.Cookie) ResponseOpt {
	return func(response *Response) {
		response.Cookies = cookie
	}
}

func ResponseWithStatusCode(statusCode int) ResponseOpt {
	return func(response *Response) {
		response.StatusCode = statusCode
	}
}

//Success 业务正常运行, 状态码200
func Success(c *gin.Context, data interface{}, args ...ResponseOpt) {
	result := NewResult(data)

	response := NewResponse(result, args...)

	baseResponse(c, response)
}

//InvalidRequest 无效入参
func InvalidRequest(c *gin.Context, message string) {
	result := NewResult(nil, ResultWithMessage(message), ResultWithCode(InvalidParams))

	response := NewResponse(result, ResponseWithStatusCode(http.StatusBadRequest))

	baseResponse(c, response)
}

//InvalidRequestWithData 与InvalidRequest功能一致, 但是可以返回数据
func InvalidRequestWithData(c *gin.Context, message string, data interface{}) {
	result := NewResult(data, ResultWithMessage(message), ResultWithCode(InvalidParams))

	response := NewResponse(result, ResponseWithStatusCode(http.StatusBadRequest))

	baseResponse(c, response)
}

//Fail 请求正常, 但是业务流不正常
func Fail(c *gin.Context, code int, message string, args ...ResponseOpt) {
	result := NewResult(nil, ResultWithCode(code), ResultWithMessage(message))

	response := NewResponse(result, args...)

	baseResponse(c, response)
}

//FailWithData 功能与Fail一致, 只不过FailWithData可以返回数据
func FailWithData(c *gin.Context, code int, message string, data interface{}, args ...ResponseOpt) {
	result := NewResult(data, ResultWithCode(code), ResultWithMessage(message))

	response := NewResponse(result, args...)

	baseResponse(c, response)
}

//RespCustomization 自定义返回
func RespCustomization(c *gin.Context, statusCode int, code int, message string, data interface{}) {

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
