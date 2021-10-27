package demo

import (
	"ginLearn/client/demo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJson(c *gin.Context) {
	var userlogin demo.UserLogin

	if err := c.ShouldBindJSON(&userlogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if userlogin.Name != "Tom" || userlogin.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
