package util

import (
	"errors"
	"fmt"
	"ginLearn/enum"
	"ginLearn/log"
	"ginLearn/util/strutils"
	"gorm.io/gorm"
	"strings"
	"time"
)

type DBClient struct {
	dbImpl DBInterface
	db     *gorm.DB
}

type DBClientOption func(client *DBClient)

func NewDBClient(args ...DBClientOption) (client *DBClient) {
	client = &DBClient{}
	for _, arg := range args {
		arg(client)
	}
	return client
}

func DBClientWithDB(db *gorm.DB) DBClientOption {
	return func(client *DBClient) {
		client.db = db
	}
}

func DBClientWithDBImpl(dbImpl DBInterface) DBClientOption {
	return func(client *DBClient) {
		client.dbImpl = dbImpl
	}
}

//DBConfig 数据库配置
type DBConfig struct {
	Host             string           `yaml:"host" json:"host"`                         //DB IP
	Port             int              `yaml:"port" json:"port"`                         //DB 端口
	User             string           `yaml:"user" json:"user"`                         //用户名
	Password         string           `yaml:"password" json:"password"`                 //密码
	DBType           string           `yaml:"dbType" json:"dbType"`                     //数据库类型
	Database         string           `yaml:"database" json:"database"`                 //数据库名
	Location         string           `yaml:"location" json:"location"`                 //时区
	TablePrefix      string           `yaml:"tablePrefix" json:"tablePrefix"`           //表前缀
	SingularTable    bool             `yaml:"singularTable" json:"singularTable"`       //表复数禁用
	CreateBatchSize  int              `yaml:"createBatchSize" json:"createBatchSize"`   //批量创建数量
	EnableRawSQL     bool             `yaml:"enableRawSql" json:"enableRawSql"`         //打印原生SQL
	MaxFreeConnCount int              `yaml:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int              `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	FreeMaxLifetime  time.Duration    `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   //闲置连接最大存活时间, 单位: 分钟
	TimeFunc         func() time.Time //设置当前时间函数
}

//Dsn 连接URL
func (c *DBConfig) Dsn() string {
	uri := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s`,
		c.User, c.Password, c.Host, c.Port, c.Database, c.Location)
	return uri
}

type DBConfigOption func(*DBConfig)

func NewDBConfig(args ...DBConfigOption) (dbConfig *DBConfig) {
	dbConfig = &DBConfig{}

	for _, arg := range args {
		arg(dbConfig)
	}

	return dbConfig
}

type Where struct {
	//Column 字段名
	Column string
	//Value 值
	Value interface{}
	//符号
	Sign string
}

func NewWhere(column string, value interface{}, sign string) *Where {
	return &Where{Column: column, Value: value, Sign: sign}
}

func NewEqWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.Eq}
}

func NewNotEqWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.NotEq}
}

func NewGtWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.Gt}
}

func NewGteWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.Gte}
}

func NewLteWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.Lte}
}

func NewLtWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.Lt}
}

func NewInWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.In}
}

func NewNotInWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.NotIn}
}

func NewLikeWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: enum.Like}
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

func NewOrder(column string, args ...OrderOpt) (order *Order) {
	order = &Order{Column: column, Sort: enum.ASC}
	for _, arg := range args {
		arg(order)
	}

	return order
}

type OrderOpt func(order *Order)

func OrderWithSort(sort string) OrderOpt {
	return func(order *Order) {
		order.Sort = sort
	}
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

//DBInterface 不同的数据库需要实现的接口
type DBInterface interface {
	initDB(config *DBConfig) (*gorm.DB, error)

	close(db *gorm.DB) error
}

//ParseDBType 解析使用数据库类型
func ParseDBType(dbType string) DBInterface {
	switch strings.ToLower(dbType) {
	case "mysql":
		return NewMysql()
	case "postgresql":
		panic("Unsupported postgresql")
	case "sqlite":
		panic("Unsupported sqlite")
	default:
		panic("Unsupported db type")
	}
}

//InitDB 初始化DB
func InitDB(dbConfig *DBConfig) (client *DBClient, err error) {
	dbImpl := ParseDBType(dbConfig.DBType)

	gormDB, err := dbImpl.initDB(dbConfig)

	if err != nil {
		return nil, err
	}

	client = NewDBClient(DBClientWithDB(gormDB), DBClientWithDBImpl(dbImpl))

	return client, nil
}

//Close 关闭数据库连接
func (c *DBClient) Close() error {
	if c.dbImpl == nil {
		return nil
	}

	return c.dbImpl.close(c.db)
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

	orderStr := condition.OrderStr()
	if strutils.IsNotEmptyStr(orderStr) {
		db = db.Order(orderStr)
	}

	limit := condition.Limit()
	if limit >= 0 {
		db = db.Limit(limit)
	}

	offset := condition.Offset()
	if offset >= 0 {
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
			query, args := where.ForWhere()
			db = db.Where(query, args)
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
