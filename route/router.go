package route

import (
	"ginLearn/api/user/handlers"
	"github.com/gin-gonic/gin"
)

func GetRoute() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	user := api.Group("/user")

	{
		user.GET("/", handlers.UserBasic)
	}

	return router
}
