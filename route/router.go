package route

import (
	projectHandlers "ginLearn/api/project/handlers"
	userHandlers "ginLearn/api/user/handlers"
	"github.com/gin-gonic/gin"
)

func BindRoute(router *gin.Engine) *gin.Engine {

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
