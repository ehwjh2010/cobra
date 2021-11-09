package controller

import (
	"fmt"
	"ginLearn/conf"
	"ginLearn/src/dao"
	"ginLearn/src/dao/model"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, conf.Conf)

	utils.Info("你好")
}

//AddRecord 添加记录
func AddRecord(c *gin.Context) {

	product := model.Product{
		Name:       "appleNormalStr",
		Price:      30,
		TotalCount: 10000,
		Brand:      utils.NewStr("oooooo"),
	}

	err := dao.DBClient.AddRecord(&product)

	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(10000), utils.ResponseWithMessage("Insert failed!"))
		return
	}

	utils.Success(c, product)
}

func UpdateRecord(c *gin.Context) {
	product := model.NewProduct()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1000))
		return
	}

	//product.TotalCount = 99
	//product.Price = 9900
	//err = dao.DBClient.UpdateById(product.TableName(), int64(id), product)

	product.ID = int64(id)
	product.TotalCount = 10
	product.Price = 156
	err = dao.DBClient.UpdateByStruct(product)

	if err != nil {
		utils.Fail(
			c,
			utils.ResponseWithCode(2000),
			utils.ResponseWithMessage(fmt.Sprintf("Update failed, %+v\n", err)))
		return
	}

	utils.Success(c, product)
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

	exist, err := dao.DBClient.QueryById(int64(idInt), product)
	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1000))
		return
	}

	if !exist {
		utils.Success(c, nil)
		return
	}

	utils.Success(c, product)
}

//QueryByCache 查缓存
func QueryByCache(c *gin.Context) {
	name := c.Param("name")

	nameValue, err := dao.CacheClient.GetString(name)

	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1000))
		return
	}

	utils.Success(c, map[string]utils.NullString{"name": nameValue})

}

//SetJob 查缓存
func SetJob(c *gin.Context) {
	job := c.Param("job")

	err := dao.CacheClient.Set("job", job, 300)

	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1000))
		return
	}

	utils.Success(c, map[string]bool{"ok": true})

}

//GetJob 查缓存
func GetJob(c *gin.Context) {
	job, err := dao.CacheClient.GetString("job")

	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1000))
		return
	}

	utils.Success(c, map[string]utils.NullString{"job": job})

}
