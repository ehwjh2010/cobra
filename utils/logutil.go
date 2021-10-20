package utils

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitLog(filepath string) {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create(filepath)
	gin.DefaultWriter = io.MultiWriter(f)
}
