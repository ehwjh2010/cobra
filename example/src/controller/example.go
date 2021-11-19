package controller

//import (
//	"cobra/conf"
//	"cobra/enum"
//	dao2 "cobra/example/src/dao"
//	model2 "cobra/example/src/dao/model"
//	"cobra/log"
//	"cobra/types"
//	"cobra/util"
//	"github.com/gin-gonic/gin"
//	"strconv"
//)
//
//func GetProjectConfig(c *gin.Context) {
//	c.JSON(200, conf.Conf)
//
//	log.Info("你好")
//}
//
////AddRecord 添加记录
//func AddRecord(c *gin.Context) {
//
//	product := model2.Product{
//		Name:       "appleNormalStr",
//		Price:      30,
//		TotalCount: 10000,
//		Brand:      types.NewStr("oooooo"),
//	}
//
//	err := dao2.DBClient.AddRecord(&product)
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(10000), util.ResultWithMessage("Insert failed!"))
//		return
//	}
//
//	util.Success(c, product)
//}
//
//func UpdateRecord(c *gin.Context) {
//	product := model2.NewProduct()
//
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	product.TotalCount = 99
//	product.Price = 9900
//	err = dao2.DBClient.UpdateById(product.TableName(), int64(id), product)
//
//	if err != nil {
//		//util.Fail(
//		//	c,
//		//	util.ResultWithCode(2000),
//		//	util.ResultWithMessage(fmt.Sprintf("Update failed, %+v\n", err)))
//		return
//	}
//
//	util.Success(c, product)
//}
//
////QueryById 通过ID查询
//func QueryById(c *gin.Context) {
//
//	id := c.Param("id")
//
//	idInt, err := strconv.Atoi(id)
//
//	if err != nil {
//		util.InvalidRequest(c, "Id must be int")
//		return
//	}
//
//	product := model2.NewProduct()
//
//	exist, err := dao2.DBClient.QueryById(int64(idInt), product)
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	if !exist {
//		util.Success(c, nil)
//		return
//	}
//
//	util.Success(c, product)
//}
//
//func QueryByCond(c *gin.Context) {
//	names := c.QueryArray("name")
//
//	page, _ := strconv.Atoi(c.Query("page"))
//	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
//	cond := util.NewQueryCondition()
//
//	cond.SetPage(page).SetPageSize(pageSize).AddSort(util.NewOrder("price", util.OrderWithSort(enum.DESC))).AddSort(util.NewOrder("id"))
//
//	cond.SetTotalCount(true)
//
//	cond.AddWhere(util.NewNotEqWhere("total_count", 90))
//
//	if len(names) > 0 {
//		cond.AddWhere(util.NewInWhere("name", names))
//	}
//
//	var products []*model2.Product
//
//	count, _ := dao2.DBClient.Query(model2.NewProduct().TableName(), cond, &products)
//
//	util.Success(c, map[string]interface{}{
//		"totalCount": count,
//		"products":   &products,
//		"page":       page,
//		"pageSize":   pageSize,
//	})
//}
//
////QueryCountByCond 查询数量
//func QueryCountByCond(c *gin.Context) {
//	product := model2.NewProduct()
//
//	cond := util.NewQueryCondition()
//
//	cond.AddWhere(util.NewEqWhere("total_count", 10))
//	cond.AddWhere(util.NewEqWhere("price", 30))
//
//	count, err := dao2.DBClient.QueryCount(product.TableName(), cond)
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(991111))
//		return
//	}
//
//	util.Success(c, map[string]int64{"count": count})
//}
//
////QueryByCache 查缓存
//func QueryByCache(c *gin.Context) {
//	name := c.Param("name")
//
//	nameValue, err := dao2.CacheClient.GetString(name)
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	util.Success(c, map[string]types.NullString{"name": nameValue})
//
//}
//
////SetJob 查缓存
//func SetJob(c *gin.Context) {
//	job := c.Param("job")
//
//	err := dao2.CacheClient.Set("job", job, 300)
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	util.Success(c, map[string]bool{"ok": true})
//
//}
//
////GetJob 查缓存
//func GetJob(c *gin.Context) {
//	job, err := dao2.CacheClient.GetString("job")
//
//	if err != nil {
//		//util.Fail(c, util.ResultWithCode(1000))
//		return
//	}
//
//	util.Success(c, map[string]types.NullString{"job": job})
//
//}
