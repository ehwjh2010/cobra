package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

//RecoveryMiddleWare 全局recovery中间件
// TODO Custom Recovery behavior
func RecoveryMiddleWare() gin.HandlerFunc {
	return gin.Recovery()
}

func init() {
	AddMiddleWares(RecoveryMiddleWare)
	log.Println("Add recovery middleware.")
}
