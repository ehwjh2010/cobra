package handlers

import (
	"ginLearn/conf"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, conf.Conf)

	utils.Log.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Info("记录一下日志", "Info")
	//Error级别的日志
	utils.Log.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Error("记录一下日志", "Error")
	//Warn级别的日志
	utils.Log.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Warn("记录一下日志", "Warn")
	//Debug级别的日志
	utils.Log.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Debug("记录一下日志", "Debug")
}
