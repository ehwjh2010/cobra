package utils

import (
	"fmt"
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

type DBConfig struct {
	Host             string        `yaml:"host" json:"host"`                         //DB IP
	Port             int           `yaml:"port" json:"port"`                         //DB 端口
	User             string        `yaml:"user" json:"user"`                         //用户名
	Password         string        `yaml:"password" json:"password"`                 //密码
	DBType           string        `yaml:"dbType" json:"dbType"`                     //数据库类型
	Database         string        `yaml:"database" json:"database"`                 //数据库名
	Location         string        `yaml:"location" json:"location"`                 //时区
	TablePrefix      string        `yaml:"tablePrefix" json:"tablePrefix"`           //表前缀
	SingularTable    bool          `yaml:"singularTable" json:"singularTable"`       //表复数禁用
	CreateBatchSize  int           `yaml:"createBatchSize" json:"createBatchSize"`   //批量创建数量
	EnableRawSQL     bool          `yaml:"enableRawSql" json:"enableRawSql"`         //打印原生SQL
	MaxFreeConnCount int           `yaml:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int           `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	FreeMaxLifetime  time.Duration `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   //闲置连接最大存活时间, 单位: 分钟
}

func (c *DBConfig) Dsn() string {
	uri := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s`,
		c.User, c.Password, c.Host, c.Port, c.Database, c.Location)
	return uri
}

type DBConfigOption func(*DBConfig)

const (
	Host             = "127.0.0.1"
	Port             = 3306
	User             = "root"
	Password         = "123456"
	Database         = "mysql"
	Location         = "UTC"
	TablePrefix      = ""
	SingularTable    = true
	CreateBatchSize  = 1000
	EnableRawSQL     = false
	MaxFreeConnCount = 3
	MaxOpenConnCount = 10
	FreeMaxLifetime  = 30
)

func NewDBConfig(args ...DBConfigOption) (dbConfig *DBConfig) {
	dbConfig = &DBConfig{
		Host:             Host,
		Port:             Port,
		User:             User,
		Password:         Password,
		Database:         Database,
		Location:         Location,
		TablePrefix:      TablePrefix,
		SingularTable:    SingularTable,
		CreateBatchSize:  CreateBatchSize,
		EnableRawSQL:     EnableRawSQL,
		MaxFreeConnCount: MaxFreeConnCount,
		MaxOpenConnCount: MaxOpenConnCount,
		FreeMaxLifetime:  FreeMaxLifetime,
	}

	for _, arg := range args {
		arg(dbConfig)
	}

	return dbConfig
}

func DBConfigWithHost(host string) DBConfigOption {
	return func(config *DBConfig) {
		config.Host = host
	}
}

func DBConfigWithPort(port int) DBConfigOption {
	return func(config *DBConfig) {
		config.Port = port
	}
}

func DBConfigWithUser(user string) DBConfigOption {
	return func(config *DBConfig) {
		config.User = user
	}
}

func DBConfigWithPwd(password string) DBConfigOption {
	return func(config *DBConfig) {
		config.Password = password
	}
}

func DBConfigWithDatabase(database string) DBConfigOption {
	return func(config *DBConfig) {
		config.Database = database
	}
}

func DBConfigWithLocation(location string) DBConfigOption {
	return func(config *DBConfig) {
		config.Location = location
	}
}

func DBConfigWithTablePrefix(tablePrefix string) DBConfigOption {
	return func(config *DBConfig) {
		config.TablePrefix = tablePrefix
	}
}

func DBConfigWithSingularTable(singularTable bool) DBConfigOption {
	return func(config *DBConfig) {
		config.SingularTable = singularTable
	}
}

func DBConfigWithCreateBatchSize(createBatchSize int) DBConfigOption {
	return func(config *DBConfig) {
		config.CreateBatchSize = createBatchSize
	}
}

func DBConfigWithEnableRawSQL(enableRawSql bool) DBConfigOption {
	return func(config *DBConfig) {
		config.EnableRawSQL = enableRawSql
	}
}

func DBConfigWithMaxFreeConnCount(maxFreeConnCount int) DBConfigOption {
	return func(config *DBConfig) {
		config.MaxFreeConnCount = maxFreeConnCount
	}
}

func DBConfigWithMaxOpenConnCount(maxOpenConnCount int) DBConfigOption {
	return func(config *DBConfig) {
		config.MaxOpenConnCount = maxOpenConnCount
	}
}

func DBConfigWithFreeMaxLifetime(freeMaxLifetime time.Duration) DBConfigOption {
	return func(config *DBConfig) {
		config.FreeMaxLifetime = freeMaxLifetime
	}
}

//DBInterface 不同的数据库需要实现的接口
type DBInterface interface {
	initDB(config *DBConfig) (*gorm.DB, error)

	close(db *gorm.DB) error
}

//ParseDBType 解析使用数据库类型
func ParseDBType(dbType string) DBInterface {
	switch dbType {
	case "mysql":
		return NewMysql()
	case "postgresql":
		panic("Unsupported db type")
	case "sqlite":
		panic("Unsupported db type")
	default:
		panic("Unsupported db type")
	}
}

//InitDB 初始化DB
func InitDB(dbConfig *DBConfig, client *DBClient) error {
	dbType := strings.ToLower(dbConfig.DBType)

	dbImpl := ParseDBType(dbType)

	gormDB, err := dbImpl.initDB(dbConfig)

	if err != nil {
		return err
	}

	client = NewDBClient(DBClientWithDB(gormDB), DBClientWithDBImpl(dbImpl))

	return nil
}

//Close 关闭数据库连接
func (c *DBClient) Close() error {
	if c.dbImpl == nil {
		return nil
	}

	return c.dbImpl.close(c.db)
}
