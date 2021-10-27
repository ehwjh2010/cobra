package user

import "github.com/gin-gonic/gin"

func UserBasic(c *gin.Context) {
	c.JSON(200, gin.H{
		"name":    "Tom",
		"company": "Google",
		"job":     "Project Manager",
	})
}
