package utils

import (
	"github.com/gin-gonic/gin"
	"io"
)

func InitLog(logDir string) {
	gin.DisableConsoleColor()
	// Logging to a file.
	logFilePath := PathJoin(logDir, "application.log")
	f, _ := OpenFileWithAppend(logFilePath)
	gin.DefaultWriter = io.MultiWriter(f)
}
