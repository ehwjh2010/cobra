package main

import (
	projectHandlers "ginLearn/src/controller/project"
	"github.com/gin-gonic/gin"
)

func BindRoute(handler *gin.Engine) {

	api := handler.Group("/api")

	config := api.Group("/project")

	{
		config.GET("/config", projectHandlers.GetProjectConfig)
	}

}
