package rdb

import (
	"errors"
	"fmt"
	"github.com/ehwjh2010/viper/client/enums"
	"github.com/ehwjh2010/viper/client/settings"
	"github.com/ehwjh2010/viper/helper/str"
	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/routine"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

const (
	//ASC 正序
	ASC = "asc"
	//DESC 倒序
	DESC = "desc"
)

const (
	Eq    = "="
	Neq   = "!="
	Gte   = ">="
	Gt    = ">"
	Lte   = "<="
	Lt    = "<"
	In    = "in"
	NotIn = "not in"
	Like  = "like"
)

type (
	DBClient struct {
		db *gorm.DB
		//rawConfig 数据库配置
		rawConfig settings.DB
		//pCount 心跳连续失败次数
		pCount int
		//rCount 重连连续失败次数
		rCount int
		//DBType 数据库类型
		DBType enums.DBType
	}

	Where struct {
		//Column 字段名
		Column string
		//Value 值
		Value interface{}
		//Sign 符号
		Sign string
		//Or 添加
		Ors []*Where
	}
)

func NewDBClient(db *gorm.DB, dbType enums.DBType, rawConfig settings.DB) (client *DBClient) {
	client = &DBClient{
		db:        db,
		DBType:    dbType,
		rawConfig: rawConfig,
	}

	return client
}

func (c *DBClient) RawConfig() settings.DB {
	return c.rawConfig
}

//Heartbeat 检测心跳
func (c *DBClient) Heartbeat() error {
	db, err := c.db.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

//WatchHeartbeat 监测心跳和重连
func (c *DBClient) WatchHeartbeat() {
	//TODO 重连逻辑接口化
	fn := func() {
		waitFlag := true
		for {
			if waitFlag {
				<-time.After(enums.ThreeSecDur)
			}

			//重连失败次数大于0, 直接重连
			if c.rCount > 0 {
				//重连次数过多, 休眠1秒后重连
				if c.rCount >= 3 {
					<-time.After(enums.OneSecDur)
				}
				if ok, _ := c.replaceDB(); ok {
					c.rCount = 0
					c.pCount = 0
					waitFlag = true
				} else {
					c.rCount++
					c.pCount++
					waitFlag = false
				}
				continue
			}

			//心跳检测
			if c.Heartbeat() != nil {
				c.pCount++
				//心跳连续3次失败, 触发重连
				if c.pCount >= 3 {
					if ok, _ := c.replaceDB(); ok {
						c.rCount = 0
						c.pCount = 0
						waitFlag = true
					} else {
						c.rCount++
						waitFlag = false
					}
				}
			} else {
				c.rCount = 0
				c.pCount = 0
				waitFlag = true
			}
		}
	}

	//优先使用协程池监听, 如果没有启用协程池, 使用原生协程监听
	err := routine.AddTask(fn)
	if err != nil {
		if errors.Is(err, routine.NoEnableRoutinePool) {
			go fn()
		} else {
			log.Warn("watch heartbeat failed")
		}
	}
}

//replaceDB 替换client内部的db对象
func (c *DBClient) replaceDB() (bool, error) {
	newDB, err := InitDBWithGorm(c.rawConfig, c.DBType)
	if err != nil {
		log.Error("reconnect db failed!", zap.Int("reconnectCount", c.rCount), zap.Error(err))
		return false, err
	}

	//关闭之前的连接
	c.Close()
	c.db = newDB
	log.Info("reconnect db success")
	return true, nil
}

//TODO Where 指定字段, 连表查询, 聚合查询, GROUP BY, HAVING, DISTINCT, COUNT, JOIN

//Or 添加Or条件
func (where *Where) Or(w *Where) *Where {
	if w != nil {
		where.Ors = append(where.Ors, w)
	}

	return where
}

//NewWhere 设置查询条件
func NewWhere(column string, value interface{}, sign string) *Where {
	return &Where{Column: column, Value: value, Sign: sign}
}

//NewEqWhere =
func NewEqWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Eq}
}

//NewNotEqWhere !=
func NewNotEqWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Neq}
}

//NewGtWhere >
func NewGtWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Gt}
}

//NewGteWhere >=
func NewGteWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Gte}
}

//NewLteWhere <=
func NewLteWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Lte}
}

//NewLtWhere <
func NewLtWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: Lt}
}

//NewInWhere in
func NewInWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: In}
}

//NewNotInWhere not in
func NewNotInWhere(column string, value interface{}) *Where {
	return &Where{Column: column, Value: value, Sign: NotIn}
}

//NewLikeWhere 模糊查询 %demo%
func NewLikeWhere(column string, value string) *Where {
	return &Where{Column: column, Value: `%` + value + `%`, Sign: Like}
}

//NewLeftLikeWhere 模糊查询 %demo
func NewLeftLikeWhere(column string, value string) *Where {
	return &Where{Column: column, Value: `%` + value, Sign: Like}
}

//NewRightLikeWhere 模糊查询 demo%
func NewRightLikeWhere(column string, value string) *Where {
	return &Where{Column: column, Value: value + `%`, Sign: Like}
}

//internal 获取SQL
func (where *Where) internal() (pattern string, value interface{}) {
	pattern = where.Column + " " + where.Sign + " " + "?"
	value = where.Value
	return
}

//tranArgs 暂不支持ors里面嵌套or
func (where *Where) tranOr(valMap map[string]interface{}, dive bool) (args []string) {

	if where.Ors == nil {
		return
	}

	// 暂不支持ors里面嵌套or
	for idx, ele := range where.Ors {
		placeholder := "@" + ele.Column + strconv.Itoa(idx)
		arg := ele.Column + " " + ele.Sign + " " + placeholder
		valMap[placeholder] = ele.Value
		args = append(args, arg)
	}

	return
}

type Order struct {
	//Column 字段名
	Column string
	//Sort 排序, 用 ASC, DESC 常量
	Sort string
}

func (o Order) String() string {
	return "order by " + o.Column + " " + o.Sort
}

//NewOrder 正排序 order by asc
func NewOrder(column string) (order *Order) {
	order = &Order{Column: column, Sort: ASC}
	return order
}

//NewDescOrder 逆排序 order by desc
func NewDescOrder(column string) (order *Order) {
	order = &Order{Column: column, Sort: DESC}
	return order
}

//description 获取排序SQL
func (o *Order) description() (description string) {
	description = o.Column + " " + o.Sort
	return description
}

//QueryCondition 查询条件
type QueryCondition struct {
	//Where 查询条件
	Where []*Where

	//Page 页数, 从1开始
	Page int

	//PageSize 每页数量, 必须大于0
	PageSize int

	//Sort 排序
	Sort []*Order

	//TotalCount 是否查询总数量
	TotalCount bool

	//Offset 偏移量
	Offset int

	//Limit  查询数量
	Limit int
}

func NewQueryCondition() *QueryCondition {
	cond := &QueryCondition{}

	return cond
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

//orderStr 获取Order排序
func (qc *QueryCondition) orderStr() string {
	if qc.Sort == nil {
		return ""
	}

	tmp := make([]string, len(qc.Sort))
	for index, item := range qc.Sort {
		tmp[index] = item.description()
	}

	result := strings.Join(tmp, ", ")

	return result
}

//GetOffset 获取偏移量
func (qc *QueryCondition) GetOffset() (offset int) {

	if qc.Offset > 0 {
		return qc.Offset
	}

	if o := qc.getOffsetByPage(); o > 0 {
		return o
	}

	return 0
}

//GetLimit 获取偏移量
func (qc *QueryCondition) GetLimit() (limit int) {

	if qc.Limit > 0 {
		return qc.Limit
	}

	if l := qc.getLimitByPage(); l > 0 {
		return l
	}

	return 0
}

//getOffsetByPage 获取偏移量
func (qc *QueryCondition) getOffsetByPage() (offset int) {
	if qc.Page < 1 {
		return 0
	}

	offset = (qc.Page - 1) * qc.PageSize

	return offset
}

//getLimitByPage 获取Limit
func (qc *QueryCondition) getLimitByPage() (limit int) {
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
		log.Err("operate db occur err", tx.Error)
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
	whereLoop:
		for _, where := range condition.Where {
			query, args := where.internal()
			db = db.Where(query, args)
			if where.Ors == nil {
				continue whereLoop
			}

			for _, w := range where.Ors {
				q, ag := w.internal()
				db = db.Or(q, ag)
			}
		}
	}

	if condition.TotalCount {
		db = db.Count(&totalCount)
	}

	if orderStr := condition.orderStr(); str.IsNotEmpty(orderStr) {
		db = db.Order(orderStr)
	}

	if limit := condition.GetLimit(); limit >= 0 {
		db = db.Limit(limit)
	}

	if offset := condition.GetOffset(); offset >= 0 {
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

	if len(condition.Where) > 0 {
	whereLoop:
		for _, where := range condition.Where {
			query, arg := where.internal()
			db = db.Where(query, arg)
			if where.Ors == nil {
				continue whereLoop
			}

			for _, w := range where.Ors {
				q, ag := w.internal()
				db = db.Or(q, ag)
			}
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

//UpdateRecordNoCond 无条件更新记录
//tableName 表名
//dstValue,  struct时, 只会更新非零字段; map 时, 根据 `map` 更新属性
func (c *DBClient) UpdateRecordNoCond(tableName string, dstValue interface{}) error {
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

//GetDB 获取原生DB对象
func (c *DBClient) GetDB() *gorm.DB {
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
		log.Err("Close conn; get db failed!", err)
		return err
	}

	err = s.Close()

	if err != nil {
		log.Error("Close mysql failed!")
	} else {
		log.Debug("Close mysql success!")
	}

	return err
}
