package middleware

import (
	"github.com/gin-gonic/gin"
)

type MiddleWareFunc func() gin.HandlerFunc

var middlewares []MiddleWareFunc

//UseMiddleWare 使用全局中间件
func UseMiddleWare(handler *gin.Engine) {

	if middlewares == nil {
		return
	}

	for _, middleware := range middlewares {
		handler.Use(middleware())
	}
}

//addMiddleWares 添加全局中间件
func addMiddleWares(mids ...MiddleWareFunc) {
	for _, middleware := range mids {
		middlewares = append(middlewares, middleware)
	}
}
