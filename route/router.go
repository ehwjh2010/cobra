package route

import (
	projectHandlers "ginLearn/api/project/handlers"
	userHandlers "ginLearn/api/user/handlers"
	"github.com/gin-gonic/gin"
)

func GetRoute() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	user := api.Group("/user")

	{
		user.GET("", userHandlers.UserBasic)
	}

	config := api.Group("/project")

	{
		config.GET("/config", projectHandlers.GetProjectConfig)
	}

	return router
}
