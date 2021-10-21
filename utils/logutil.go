package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func InitLog(logDir string, application string) {
	gin.DisableConsoleColor()

	fileName := fmt.Sprintf(`%s.log`, application)
	// Logging to a file.
	logFilePath := PathJoin(logDir, fileName)
	f, _ := OpenFileWithAppend(logFilePath)
	gin.DefaultWriter = io.MultiWriter(f)
}
