package demo

import (
	"fmt"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

//MethodGetDemo 查询字符传
func MethodGetDemo(c *gin.Context) {
	name := c.Query("name")
	gender := c.DefaultQuery("gender", "unknown")
	age := c.Query("age")
	limit := c.Query("limit")
	offset := c.DefaultQuery("offset", "0")

	c.JSON(http.StatusOK, gin.H{
		"name":   name,
		"age":    age,
		"gender": gender,
		"limit":  limit,
		"offset": offset,
	})
}

//MethodPathDemo 路径参数
func MethodPathDemo(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")

	fmt.Println(c.FullPath())

	//c.JSON(http.StatusOK, gin.H{
	//	"name": name,
	//	"action": action,
	//})

	c.String(http.StatusOK, name+" is "+action)
}

//MethodUploadDemo 文件上传
func MethodUploadDemo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Infof("Get file failed! err: %v", err)
	}

	filename := file.Filename

	currentDir, _ := os.Getwd()

	c.SaveUploadedFile(file, utils.PathJoin(currentDir, "logs", filename))

	c.String(http.StatusOK, filename)
}

//MethodJson 请求体读取
func MethodJson(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		utils.Infof("Read body json failed!")
	}

	c.String(http.StatusOK, string(data))
}
