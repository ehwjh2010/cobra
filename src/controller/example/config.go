package example

import (
	"ginLearn/src/dao"
	"ginLearn/src/dao/model"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, "1111")

	utils.Info("你好")
}

//QueryById 通过ID查询
func QueryById(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		utils.InvalidRequest(c, "Id must be int")
		return
	}

	product := model.NewProduct()

	_, err = dao.DBClient.QueryById(int64(idInt), product)
	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1000))
		return
	}

	utils.Success(c, product)
}
