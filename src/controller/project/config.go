package project

import (
	"ginLearn/src/configure"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, configure.Conf)

	utils.Info()
}
