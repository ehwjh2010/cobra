package response

import (
	"github.com/ehwjh2010/viper/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Success 业务正常运行, 状态码200
func Success(c *gin.Context, data interface{}, args ...types.RespOpt) {
	result := types.NewResult(data)

	response := types.NewResp(result, args...)

	baseResponse(c, response)
}

//InvalidRequest 无效入参, 状态码400
func InvalidRequest(c *gin.Context, message string) {
	result := types.NewErrResp(types.InvalidParams, message)

	response := types.NewResp(result, types.RespWithStatus(http.StatusBadRequest))

	baseResponse(c, response)
}

//InvalidRequestWithData 与InvalidRequest功能一致, 但是可以返回数据
func InvalidRequestWithData(c *gin.Context, message string, data interface{}) {
	result := types.NewResult(data, types.ResultWithMessage(message), types.ResultWithCode(types.InvalidParams))

	response := types.NewResp(result, types.RespWithStatus(http.StatusBadRequest))

	baseResponse(c, response)
}

//Fail 请求正常, 状态码200, 但是业务流不正常
func Fail(c *gin.Context, code int, message string, args ...types.RespOpt) {
	result := types.NewErrResp(code, message)

	response := types.NewResp(result, args...)

	baseResponse(c, response)
}

func FailWithResult(c *gin.Context, result types.Result, args ...types.RespOpt) {
	response := types.NewResp(result, args...)

	baseResponse(c, response)
}

//FailWithData 功能与Fail一致, 只不过FailWithData可以返回数据
func FailWithData(c *gin.Context, code int, message string, data interface{}, args ...types.RespOpt) {
	result := types.NewResult(data, types.ResultWithCode(code), types.ResultWithMessage(message))

	response := types.NewResp(result, args...)

	baseResponse(c, response)
}

func CustomResponse(c *gin.Context, response *types.Response) {
	baseResponse(c, response)
}

func baseResponse(c *gin.Context, response *types.Response) {
	if len(response.Header) > 0 {
		for key, value := range response.Header {
			c.Header(key, value)
		}
	}

	if response.Cookie != nil {
		for _, cookie := range response.Cookie {
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
