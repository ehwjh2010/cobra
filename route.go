package main

import (
	"ginLearn/src/controller/example"
	"github.com/gin-gonic/gin"
)

func BindRoute(handler *gin.Engine) {

	api := handler.Group("/api")

	exampleGroup := api.Group("/example")

	{
		exampleGroup.GET("/config", example.GetProjectConfig)
		exampleGroup.GET("/db/:id", example.QueryById)
		exampleGroup.GET("/db/add", example.AddRecord)
		exampleGroup.GET("/cache/:name", example.QueryByCache)
		exampleGroup.GET("/cache/set/:job", example.SetJob)
		exampleGroup.GET("/cache/get/job", example.GetJob)

	}

}
