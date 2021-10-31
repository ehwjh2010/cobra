package middleware

import (
	"github.com/gin-gonic/gin"
)

type MiddleWareFunc func() gin.HandlerFunc

var middlewares []MiddleWareFunc

//UseMiddleWares 使用全局中间件
func UseMiddleWares(server *gin.Engine) {

	if middlewares == nil {
		return
	}

	for _, middleware := range middlewares {
		server.Use(middleware())
	}
}

//AddMiddleWares 添加全局中间件
func AddMiddleWares(mids ...MiddleWareFunc) {
	for _, middleware := range mids {
		middlewares = append(middlewares, middleware)
	}
}
