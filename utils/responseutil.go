package utils

import "github.com/gin-gonic/gin"
import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseOption func(response *Response)

const (
	OK           = 0
	InvalidParam = 10
)

const UnknownMessage = "Unknown message"

var MessageMap = map[int]string{
	OK:           "Success",
	InvalidParam: "Invalid param",
}

func GetMsg(code int) string {
	value, exist := MessageMap[code]
	if exist {
		return value
	} else {
		return UnknownMessage
	}
}

func NewResponse(args ...ResponseOption) (response *Response) {
	response = &Response{
		Code:    OK,
		Message: GetMsg(OK),
		Data:    nil,
	}
	for _, arg := range args {
		arg(response)
	}

	return response
}

func ResponseWithCode(code int) ResponseOption {
	return func(response *Response) {
		response.Code = code
	}
}

func ResponseWithMessage(msg string) ResponseOption {
	return func(response *Response) {
		response.Message = msg
	}
}

func ResponseWithData(data interface{}) ResponseOption {
	return func(response *Response) {
		response.Data = data
	}
}

func NewInvalidParamResp(data interface{}) (response *Response) {
	response = NewResponse(
		ResponseWithCode(InvalidParam),
		ResponseWithMessage(GetMsg(InvalidParam)),
		ResponseWithData(data))
	return response
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(ResponseWithData(data)))
}

func InvalidRequest(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewInvalidParamResp(data))
}

func Fail(c *gin.Context, args ...ResponseOption) {
	c.JSON(http.StatusOK, NewResponse(args...))
}
