package middleware

import (
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
)

type MiddleWareFunc func() gin.HandlerFunc

var middlewares []MiddleWareFunc

func UseMiddleWares(server *gin.Engine) {

	if utils.IsNil(middlewares) {
		return
	}

	for _, middleware := range middlewares {
		server.Use(middleware())
	}
}

func AddMiddleWares(mids ...MiddleWareFunc) {
	for _, middleware := range mids {
		middlewares = append(middlewares, middleware)
	}
}

func init() {
	AddMiddleWares(LoggerToFile, gin.Recovery)
}
