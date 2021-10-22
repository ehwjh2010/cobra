package utils

import (
	"fmt"
	"ginLearn/structs/setting"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

func InitLog(conf setting.Config) {
	gin.DisableConsoleColor()
	// Logging to a file.

	dirLogPath := PathJoin(conf.LogPath, conf.Application)
	err := MakeDirs(dirLogPath)
	if err != nil {
		log.Fatalf("Access log dir failed! err: %v", err)
	}

	logFilePath := PathJoin(dirLogPath, "application.log")
	f, _ := OpenFileWithAppend(logFilePath)

	writers := []io.Writer{f}
	if conf.EnableLogConsole {
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
