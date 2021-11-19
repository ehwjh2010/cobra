package middleware

import (
	"github.com/gin-gonic/gin"
)

func UseMiddlewares(engine *gin.Engine, middlewares ...gin.HandlerFunc) {
	if middlewares == nil {
		return
	}

	for _, middleware := range middlewares {
		engine.Use(middleware)
	}
}
