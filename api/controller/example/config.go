package example

import (
	"ginLearn/api/dao"
	"ginLearn/api/dao/model"
	"ginLearn/conf"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, conf.Conf)

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

	name, err := dao.CacheClient.GetString(name)

	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1000))
		return
	}

	utils.Success(c, map[string]string{"name": name})

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

	utils.Success(c, map[string]string{"job": job})

}
