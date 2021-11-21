package middleware

import (
	"github.com/gin-gonic/gin"
)

func UseMiddlewares(engine *gin.Engine, middlewares ...gin.HandlerFunc) {

	for _, middleware := range middlewares {
		if middleware == nil {
			continue
		}
		engine.Use(middleware)
	}
}
