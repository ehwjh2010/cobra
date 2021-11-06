package project

import (
	"ginLearn/resource"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, resource.Conf)

	utils.Info("你好")
}
