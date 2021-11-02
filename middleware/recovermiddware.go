package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

//RecoveryMiddleWare 全局recovery中间件
func RecoveryMiddleWare() gin.HandlerFunc {
	return gin.Recovery()
}

func init() {
	addMiddleWares(RecoveryMiddleWare)
	log.Println("Add recovery middleware.")
}
