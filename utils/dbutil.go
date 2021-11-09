package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type DBClient struct {
	dbImpl DBInterface
	DB     *gorm.DB
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
		client.DB = db
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
		panic("Unsupported DB type")
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

	return c.dbImpl.close(c.DB)
}

//occurErr 判断是否发生报错
func (c *DBClient) occurErr(tx *gorm.DB, excludeErr ...error) bool {

	if tx.Error == nil {
		return false
	}

	txErr := tx.Error
	if excludeErr == nil {
		return true
	}

	for _, err := range excludeErr {
		if errors.Is(txErr, err) {
			return false
		}
	}

	return true
}

func (c *DBClient) check(tx *gorm.DB, excludeErr ...error) (exist bool, err error) {

	if c.occurErr(tx, excludeErr...) {
		Errorf("Query DB occur err, err: %v", tx.Error)
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
	return c.DB.AutoMigrate(pointers...)
}

//QueryById 通过主键查询
//exist 记录是否存在
//err 发生的错误
func (c *DBClient) QueryById(id int64, pointer interface{}) (exist bool, err error) {
	db := c.DB

	tx := db.Limit(1).Where("id = ?", id).Find(pointer)

	return c.check(tx)
}

//QueryByIds 通过主键查询
func (c *DBClient) QueryByIds(ids []int64, pointers interface{}) (exist bool, err error) {
	tx := c.DB.Where("id in ?", ids).Find(pointers)
	return c.check(tx)
}

//QueryByStruct 通过结构体查询, 结构体字段为零值的字段, 不会作为条件
func (c *DBClient) QueryByStruct(condition interface{}, dst interface{}) (exist bool, err error) {
	tx := c.DB.Where(condition).Find(dst)
	return c.check(tx)
}

//QueryByMap 通过Map查询
func (c *DBClient) QueryByMap(condition map[string]interface{}, dst interface{}) (exist bool, err error) {
	tx := c.DB.Where(condition).Find(dst)
	return c.check(tx)
}

//First 查询第一条记录
func (c *DBClient) First(condition interface{}, pointer interface{}) (exist bool, err error) {
	tx := c.DB.Where(condition).First(pointer)

	return c.check(tx, gorm.ErrRecordNotFound)
}

//Last 查询最后一条记录
func (c *DBClient) Last(condition interface{}, pointer interface{}) (exist bool, err error) {
	tx := c.DB.Where(condition).Last(pointer)

	return c.check(tx, gorm.ErrRecordNotFound)
}

//Exist 记录是否存在
func (c *DBClient) Exist(condition map[string]interface{}, dst interface{}) (exist bool, err error) {
	return c.First(condition, dst)
}

//AddRecord 添加记录
//data 指针
func (c *DBClient) AddRecord(data interface{}) error {
	tx := c.DB.Create(data)

	return tx.Error
}

//AddRecords 批量添加记录
func (c *DBClient) AddRecords(data interface{}, batchSize int) error {
	tx := c.DB.CreateInBatches(data, batchSize)
	return tx.Error
}

//UpdateById 根据主键更新
//data为结构体指针时, 结构体零值字段不会被更新
//data为`map`时, 更具`map`更新属性
func (c *DBClient) UpdateById(tableName string, id int64, data interface{}) error {
	tx := c.DB.Table(tableName).Where("id = ?", id).Updates(data)
	return tx.Error
}

//UpdateRecord 更新记录, condition必须包含条件, 否则会返回错误ErrMissingWhereClause,
//如果想无条件更新, 请使用updateRecordWithoutCond
//tableName  表名
//dstValue	 struct时, 只会更新非零字段; map 时, 根据 `map` 更新属性
//condition	 struct时, 只会把非零字段当做条件; map 时, 根据 `map` 设置条件
func (c *DBClient) UpdateRecord(tableName string, condition interface{}, dstValue interface{}) error {
	tx := c.DB.Table(tableName).Where(condition).Updates(dstValue)
	return tx.Error
}

//UpdateByStruct 根据主键更新, data必须为结构体指针且主键为有效值, 结构体零值字段不会被更新
func (c *DBClient) UpdateByStruct(data interface{}) error {
	tx := c.DB.Updates(data)
	return tx.Error
}

//UpdateRecordWithoutCond 无条件更新记录
//tableName 表名
//dstValue,  struct时, 只会更新非零字段; map 时, 根据 `map` 更新属性
func (c *DBClient) UpdateRecordWithoutCond(tableName string, dstValue interface{}) error {
	tx := c.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Table(tableName).Updates(dstValue)
	return tx.Error
}

//Raw 执行原生SQL
func (c *DBClient) Raw(sql string) error {
	tx := c.DB.Exec(sql)
	return tx.Error
}

//Save 保存记录, 会保存所有的字段，即使字段是零值
//ptr 必须是struct指针
func (c *DBClient) Save(ptr interface{}) error {
	tx := c.DB.Save(ptr)
	return tx.Error
}
