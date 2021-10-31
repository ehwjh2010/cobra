package project

import (
	"ginLearn/src/configure"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, configure.Conf)

	utils.InfoWithFields(logrus.Fields{"name": "JH"}, "你好")
}
