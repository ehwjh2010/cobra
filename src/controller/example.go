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

	product.TotalCount = 99
	product.Price = 9900
	err = dao.DBClient.UpdateById(product.TableName(), int64(id), product)

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

func QueryByCond(c *gin.Context) {
	names := c.QueryArray("name")

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	cond := utils.NewQueryCondition()

	cond.SetPage(page).SetPageSize(pageSize).AddSort(utils.NewOrder("price", utils.OrderWithSort(utils.DESC))).AddSort(utils.NewOrder("id"))

	cond.SetTotalCount(true)

	cond.AddWhere(utils.NewNotEqWhere("total_count", 90))

	if len(names) > 0 {
		cond.AddWhere(utils.NewInWhere("name", names))
	}

	var products []*model.Product

	count, _ := dao.DBClient.Query(model.NewProduct().TableName(), cond, &products)

	utils.Success(c, map[string]interface{}{
		"totalCount": count,
		"products":   &products,
		"page":       page,
		"pageSize":   pageSize,
	})
}

//QueryCountByCond 查询数量
func QueryCountByCond(c *gin.Context) {
	product := model.NewProduct()

	cond := utils.NewQueryCondition()

	cond.AddWhere(utils.NewEqWhere("total_count", 10))
	cond.AddWhere(utils.NewEqWhere("price", 30))

	count, err := dao.DBClient.QueryCount(product.TableName(), cond)

	if err != nil {
		utils.Fail(c, utils.ResponseWithCode(991111))
		return
	}

	utils.Success(c, map[string]int64{"count": count})
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
