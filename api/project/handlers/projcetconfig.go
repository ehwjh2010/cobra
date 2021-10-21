package handlers

import (
	"ginLearn/conf"
	"github.com/gin-gonic/gin"
)

func GetProjectConfig(c *gin.Context) {

	c.JSON(200, conf.Conf)
}
