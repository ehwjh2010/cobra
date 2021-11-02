package project

import (
	"ginLearn/resource"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, resource.Conf)

	utils.InfoWithFields(logrus.Fields{"name": "JH"}, "你好")
}
