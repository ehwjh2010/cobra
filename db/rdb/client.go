package rdb

import (
	"errors"
	"fmt"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/util/strutils"
	"gorm.io/gorm"
	"strings"
)

const (
	//ASC 正序
	ASC = "asc"
	//DESC 倒序
	DESC = "desc"
)

const (
	Eq    = "="
	Nq    = "!="
	Gte   = ">="
	Gt    = ">"
	Lte   = "<="
	Lt    = "<"
	In    = "in"
	NotIn = "not in"
	Like  = "like"
)

const (
	Mysql = iota
	Postgresql
	Sqlite
)

type (
	DBClient struct {
		db *gorm.DB
		//DbType 数据库类型
		DbType int
	}

	Where struct {
		//Column 字段名
		Column string
		//Value 值
		Value interface{}
		//Sign 符号
		Sign string
	}
)

func NewDBClient(db *gorm.DB, dbType int) (client *DBClient) {
	client = &DBClient{
		db:     db,
		DbType: dbType,
	}

	return client
}

func NewWhere(column string, value interface{}, sign string) *Where {
	return &Where{Column: column, Value: value, Sign: sign}
}

func NewEqWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Eq}
}

func NewNotEqWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Nq}
}

func NewGtWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Gt}
}

func NewGteWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Gte}
}

func NewLteWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Lte}
}

func NewLtWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Lt}
}

func NewInWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: In}
}

func NewNotInWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: NotIn}
}

//NewLikeWhere TODO 模糊查询 格式化值
func NewLikeWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Like}
}

//NewLeftLikeWhere TODO 模糊查询 格式化值
func NewLeftLikeWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Like}
}

//NewRightLikeWhere TODO 模糊查询 格式化值
func NewRightLikeWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Like}
}

//ForWhere 获取SQL
func (where *Where) ForWhere() (pattern string, value interface{}) {
	pattern = strings.Join([]string{where.Column, where.Sign, "?"}, " ")
	value = where.Value
	return
}

type Order struct {
	//Column 字段名
	Column string
	//Sort 排序, 用 ASC, DESC 常量
	Sort string
}

func NewOrder(column string) (order *Order) {
	order = &Order{Column: column, Sort: ASC}
	return order
}

func NewDescOrder(column string) (order *Order) {
	order = &Order{Column: column, Sort: DESC}
	return order
}

//Description 获取排序SQL
func (o *Order) Description() (description string) {
	description = fmt.Sprintf("%s %s", o.Column, o.Sort)
	return description
}

//QueryCondition 查询条件
type QueryCondition struct {
	//Where 查询条件, `map`时, 根据`map`设定查询条件, struct时, struct零值字段不会当做条件. 推荐使用`map`
	Where []*Where

	//Page 页数, 从1开始
	Page int

	//PageSize 每页数量, 必须大于0
	PageSize int

	//Sort 排序
	Sort []*Order

	//TotalCount 是否查询总数量
	TotalCount bool

	offset int

	limit int
}

func NewQueryCondition(args ...QCOption) *QueryCondition {
	cond := &QueryCondition{}

	for _, arg := range args {
		arg(cond)
	}

	return cond
}

type QCOption func(condition *QueryCondition)

func QCWithWhere(where []*Where) QCOption {
	return func(condition *QueryCondition) {
		condition.Where = where
	}
}

func QCWithPage(page int) QCOption {
	if page <= 0 {
		panic("Page must gt 0")
	}

	return func(condition *QueryCondition) {
		condition.Page = page
	}
}

func QCWithPageSize(pageSize int) QCOption {
	if pageSize <= 0 {
		panic("PageSize must gt 0")
	}

	return func(condition *QueryCondition) {
		condition.PageSize = pageSize
	}
}

func QCWithTotalCount(totalCount bool) QCOption {
	return func(condition *QueryCondition) {
		condition.TotalCount = totalCount
	}
}

func QCWithSort(sort []*Order) QCOption {
	return func(condition *QueryCondition) {
		condition.Sort = sort
	}
}

//AddWhere 添加条件
func (qc *QueryCondition) AddWhere(where *Where) *QueryCondition {
	if where != nil {
		qc.Where = append(qc.Where, where)
	}
	return qc
}

//AddSort 添加排序
func (qc *QueryCondition) AddSort(sort *Order) *QueryCondition {
	if sort != nil {
		qc.Sort = append(qc.Sort, sort)
	}
	return qc
}

//SetPage 设置页数
func (qc *QueryCondition) SetPage(page int) *QueryCondition {
	qc.Page = page
	return qc
}

//SetPageSize 设置每页数量
func (qc *QueryCondition) SetPageSize(pageSize int) *QueryCondition {
	qc.PageSize = pageSize
	return qc
}

//SetTotalCount 设置是否查询总数
func (qc *QueryCondition) SetTotalCount(query bool) *QueryCondition {
	qc.TotalCount = query
	return qc
}

//OrderStr 获取Order排序
func (qc *QueryCondition) OrderStr() string {
	if qc.Sort == nil {
		return ""
	}

	var tmp []string
	for _, item := range qc.Sort {
		tmp = append(tmp, item.Description())
	}

	result := strings.Join(tmp, ", ")

	return result
}

//Offset 获取偏移量
func (qc *QueryCondition) Offset() (offset int) {
	if qc.Page < 1 {
		return 0
	}

	offset = (qc.Page - 1) * qc.PageSize

	return offset
}

//Limit 获取Limit
func (qc *QueryCondition) Limit() (limit int) {
	return qc.PageSize
}

//occurErr 判断是否发生报错
func (c *DBClient) occurErr(tx *gorm.DB, excludeErr ...error) bool {

	if tx.Error == nil {
		return false
	}

	if excludeErr == nil {
		return true
	}

	txErr := tx.Error

	for _, err := range excludeErr {
		if errors.Is(txErr, err) {
			return false
		}
	}

	return true
}

func (c *DBClient) check(tx *gorm.DB, excludeErr ...error) (exist bool, err error) {

	if c.occurErr(tx, excludeErr...) {
		log.Errorf("Query db occur err, err: %v", tx.Error)
		return false, tx.Error
	}

	if tx.RowsAffected <= 0 {
		return false, nil
	}

	return true, nil
}

//Migrate 数据库迁移
//models 数据库模型
//model: client.Migrate(&Product{}, &Fruit{})
func (c *DBClient) Migrate(pointers ...interface{}) error {
	db := c.db

	return db.AutoMigrate(pointers...)
}

func (c *DBClient) QueryByPrimaryKey(pkColumnName string, pkValue, pointer interface{}) (exist bool, err error) {
	db := c.db

	pattern := fmt.Sprintf("%s = ?", pkColumnName)

	tx := db.Limit(1).Where(pattern, pkValue).Find(pointer)

	return c.check(tx)
}

//QueryById 通过主键查询
//exist 记录是否存在
//err 发生的错误
func (c *DBClient) QueryById(id int64, pointer interface{}) (exist bool, err error) {
	db := c.db

	tx := db.Limit(1).Where("id = ?", id).Find(pointer)

	return c.check(tx)
}

//QueryByIds 通过主键查询
func (c *DBClient) QueryByIds(ids []int64, pointers interface{}) (exist bool, err error) {
	db := c.db

	tx := db.Where("id in ?", ids).Find(pointers)

	return c.check(tx)
}

//Query 查询
func (c *DBClient) Query(tableName string, condition *QueryCondition, dst interface{}) (totalCount int64, err error) {
	db := c.db

	db = db.Table(tableName)

	if condition.Where != nil {
		for _, where := range condition.Where {
			query, args := where.ForWhere()
			db = db.Where(query, args)
		}
	}

	if condition.TotalCount {
		db = db.Count(&totalCount)
	}

	if orderStr := condition.OrderStr(); strutils.IsNotEmptyStr(orderStr) {
		db = db.Order(orderStr)
	}

	if limit := condition.Limit(); limit >= 0 {
		db = db.Limit(limit)
	}

	if offset := condition.Offset(); offset >= 0 {
		db = db.Offset(offset)
	}

	tx := db.Find(dst)

	_, err = c.check(tx)

	return totalCount, err

}

//QueryCount 查询数量
func (c *DBClient) QueryCount(tableName string, condition *QueryCondition) (count int64, err error) {
	db := c.db

	db = db.Table(tableName)

	if condition.Where != nil {
		for _, where := range condition.Where {
			query, arg := where.ForWhere()
			db = db.Where(query, arg)
		}
	}

	tx := db.Count(&count)

	_, err = c.check(tx)

	return count, err

}

//QueryByStruct 通过结构体查询, 结构体字段为零值的字段, 不会作为条件
func (c *DBClient) QueryByStruct(condition interface{}, dst interface{}) (exist bool, err error) {
	db := c.db

	tx := db.Where(condition).Find(dst)

	return c.check(tx)
}

//QueryByMap 通过Map查询
func (c *DBClient) QueryByMap(condition map[string]interface{}, dst interface{}) (exist bool, err error) {
	db := c.db

	tx := db.Where(condition).Find(dst)

	return c.check(tx)
}

//First 查询第一条记录
func (c *DBClient) First(condition interface{}, pointer interface{}) (exist bool, err error) {
	db := c.db

	tx := db.Where(condition).First(pointer)

	return c.check(tx, gorm.ErrRecordNotFound)
}

//Last 查询最后一条记录
func (c *DBClient) Last(condition interface{}, pointer interface{}) (exist bool, err error) {
	db := c.db

	tx := db.Where(condition).Last(pointer)

	return c.check(tx, gorm.ErrRecordNotFound)
}

//Exist 记录是否存在
func (c *DBClient) Exist(condition map[string]interface{}, dst interface{}) (exist bool, err error) {
	return c.First(condition, dst)
}

//AddRecord 添加记录
//data 结构体指针
func (c *DBClient) AddRecord(data interface{}) error {
	db := c.db

	tx := db.Create(data)

	return tx.Error
}

//AddRecords 批量添加记录
func (c *DBClient) AddRecords(data interface{}, batchSize int) error {
	db := c.db

	tx := db.CreateInBatches(data, batchSize)

	return tx.Error
}

//UpdateById 根据主键更新
//data为结构体指针时, 结构体零值字段不会被更新
//data为`map`时, 更具`map`更新属性
func (c *DBClient) UpdateById(tableName string, id int64, data interface{}) error {
	db := c.db

	tx := db.Table(tableName).Where("id = ?", id).Updates(data)

	return tx.Error
}

//UpdateRecord 更新记录, condition必须包含条件, 否则会返回错误ErrMissingWhereClause,
//如果想无条件更新, 请使用updateRecordWithoutCond
//tableName  表名
//dstValue	 struct时, 只会更新非零字段; map 时, 根据 `map` 更新属性
//condition	 struct时, 只会把非零字段当做条件; map 时, 根据 `map` 设置条件
func (c *DBClient) UpdateRecord(tableName string, condition interface{}, dstValue interface{}) error {
	db := c.db

	tx := db.Table(tableName).Where(condition).Updates(dstValue)

	return tx.Error
}

//UpdateRecordWithoutCond 无条件更新记录
//tableName 表名
//dstValue,  struct时, 只会更新非零字段; map 时, 根据 `map` 更新属性
func (c *DBClient) UpdateRecordWithoutCond(tableName string, dstValue interface{}) error {
	db := c.db

	tx := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Table(tableName).Updates(dstValue)

	return tx.Error
}

//Save 保存记录, 会保存所有的字段，即使字段是零值
//ptr 必须是struct指针
func (c *DBClient) Save(ptr interface{}) error {
	db := c.db

	tx := db.Save(ptr)

	return tx.Error
}

//DB 获取原生DB对象
func (c *DBClient) DB() *gorm.DB {
	db := c.db
	return db
}

//Close 关闭连接池
func (c *DBClient) Close() error {
	if c.db == nil {
		return nil
	}

	s, err := c.db.DB()
	if err != nil {
		log.Errorf("Close conn; get db failed!, err: %v", err)
		return err
	}

	err = s.Close()

	if err != nil {
		log.Errorl("Close mysql failed!")
	} else {
		log.Infol("Close mysql success!")
	}

	return err
}
