package example

import (
	"fmt"
	"ginLearn/client/example"
	"ginLearn/conf"
	"ginLearn/src/dao"
	"ginLearn/src/dao/model"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"strconv"
)

func GetProjectConfig(c *gin.Context) {
	c.JSON(200, conf.Conf)

	utils.Info("你好")
}

//AddRecord 添加记录
func AddRecord(c *gin.Context) {

	product := model.Product{
		Name:       "orange",
		Price:      30,
		TotalCount: 10000,
		Brand:      null.NewString("肯德基", true),
	}

	dao.DBClient.AddRecord(&product)

	fmt.Printf("product=%+v\n", product)
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

	p := example.NewProduct()
	err = utils.CopyProperty(product, p)
	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(1001))
		return
	}

	utils.Success(c, p)
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
