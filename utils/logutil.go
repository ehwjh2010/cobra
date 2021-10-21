package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func InitLog(logDir string, enableLogConsole bool) {
	gin.DisableConsoleColor()
	// Logging to a file.
	logFilePath := PathJoin(logDir, "application.log")
	f, _ := OpenFileWithAppend(logFilePath)
	writers := []io.Writer{f}
	if enableLogConsole {
		writers = append(writers, os.Stdout)
	}
	gin.DefaultWriter = io.MultiWriter(writers...)
}

//SetLogFormat 设置日志格式
func SetLogFormat(r *gin.Engine) {
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		fmt.Println(param)
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
}
