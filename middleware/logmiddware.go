package middleware

import (
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//LoggerToFile 日志中间件,
//TODO 中间件未生效
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 执行时间
		latencyTime := time.Since(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		utils.InfoF("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func init() {
	AddMiddleWares(LoggerToFile)
	log.Println("Add log middleware.")
}
