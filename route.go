package main

import (
	"ginLearn/src/controller"
	"github.com/gin-gonic/gin"
)

func BindRoute(handler *gin.Engine) {

	api := handler.Group("/api")

	exampleGroup := api.Group("/example")

	{
		exampleGroup.GET("/config", controller.GetProjectConfig)
		exampleGroup.GET("/db/:id", controller.QueryById)
		exampleGroup.GET("/db/query", controller.QueryByCond)
		exampleGroup.GET("/db/add", controller.AddRecord)
		exampleGroup.GET("/db/update/:id", controller.UpdateRecord)
		exampleGroup.GET("/cache/:name", controller.QueryByCache)
		exampleGroup.GET("/cache/set/:job", controller.SetJob)
		exampleGroup.GET("/cache/get/job", controller.GetJob)

	}

}
