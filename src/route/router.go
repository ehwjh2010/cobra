package route

import (
	projectHandlers "ginLearn/src/controller/project"
	userHandlers "ginLearn/src/controller/user/handlers"
	"github.com/gin-gonic/gin"
)

func BindRoute(server *gin.Engine) *gin.Engine {

	api := server.Group("/api")

	user := api.Group("/user")

	{
		user.GET("", userHandlers.UserBasic)
	}

	config := api.Group("/project")

	{
		config.GET("/config", projectHandlers.GetProjectConfig)
	}

	return server
}
