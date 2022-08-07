package rdb

import (
	"errors"
	"time"

	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/ehwjh2010/viper/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

const defaultCreateBatchSize = 1000

var UnsupportedDBType = errors.New("unsupported db type")

func InitDBWithGorm(dbConfig DB, dbType enums.DBType) (*gorm.DB, error) {

	var sqlLogger = logger.Silent
	if dbConfig.EnableRawSQL {
		sqlLogger = logger.Info
	}

	var createBatchSize = defaultCreateBatchSize
	if dbConfig.CreateBatchSize > 0 {
		createBatchSize = dbConfig.CreateBatchSize
	}

	dialector, err := getDialector(dbConfig.Url, dbType)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		// 打印SQL
		Logger: logger.Default.LogMode(sqlLogger),
		NamingStrategy: schema.NamingStrategy{
			// 指定表前缀
			TablePrefix: dbConfig.TablePrefix,
			// 表复数禁用
			SingularTable: dbConfig.SingularTable,
		},
		// 批量操作 每批数量
		CreateBatchSize: createBatchSize,
	})

	if err != nil {
		return nil, err
	}

	maxIdleConn := dbConfig.MaxFreeConnCount
	maxOpenConn := dbConfig.MaxOpenConnCount
	connMaxIdleTime := time.Duration(dbConfig.FreeMaxLifetime) * time.Second
	connMaxLifetime := time.Duration(dbConfig.ConnMaxLifetime) * time.Second

	// 注册集群
	if str.IsNotEmptySlice(dbConfig.Replicas) {
		writeDialector, _ := getDialector(dbConfig.Url, dbType)
		readerDialectors := make([]gorm.Dialector, len(dbConfig.Replicas))

		for index, replica := range dbConfig.Replicas {
			readerDialector, dialErr := getDialector(replica, dbType)
			if dialErr != nil {
				return nil, dialErr
			}
			readerDialectors[index] = readerDialector
		}

		// 设置读写节点
		useErr := db.Use(dbresolver.Register(
			dbresolver.Config{
				// 写节点
				Sources: []gorm.Dialector{writeDialector},
				// 读节点
				Replicas: readerDialectors,
				// sources/replicas 负载均衡策略
				Policy: dbresolver.RandomPolicy{},
			}).
			// 设置连接池中空闲连接最大数
			SetMaxIdleConns(maxIdleConn).
			// 设置打开数据库最大连接数
			SetMaxOpenConns(maxOpenConn).
			// 设置闲置连接最长存活时间
			SetConnMaxIdleTime(connMaxIdleTime).
			// 设置连接最大存活时间
			SetConnMaxLifetime(connMaxLifetime))

		if useErr != nil {
			return nil, useErr
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 设置连接池中空闲连接最大数
	sqlDB.SetMaxIdleConns(maxIdleConn)
	// 设置打开数据库最大连接数
	sqlDB.SetMaxOpenConns(maxOpenConn)
	// 设置闲置连接最长存活时间
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime)
	// 设置连接最大存活时间
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	// 由于gorm.processor是私有的, 无法定义函数, 有了以下重复代码
	if len(dbConfig.CreateCallbacks) > 0 {
		for _, item := range dbConfig.CreateCallbacks {
			switch item.OpType {
			case enums.Register:
				switch item.When {
				case enums.Before:
					if e := db.Callback().Create().Before(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				case enums.After:
					if e := db.Callback().Create().After(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				default:
					if e := db.Callback().Create().Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				}
			case enums.Replace:
				if e := db.Callback().Create().Replace(item.Name, item.Action); e != nil {
					return nil, e
				}
			case enums.Remove:
				if e := db.Callback().Create().Remove(item.Name); e != nil {
					return nil, e
				}
			}
		}
	}

	if len(dbConfig.UpdateCallbacks) > 0 {
		for _, item := range dbConfig.UpdateCallbacks {
			switch item.OpType {
			case enums.Register:
				switch item.When {
				case enums.Before:
					if e := db.Callback().Update().Before(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				case enums.After:
					if e := db.Callback().Update().After(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				default:
					if e := db.Callback().Update().Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				}
			case enums.Replace:
				if e := db.Callback().Update().Replace(item.Name, item.Action); e != nil {
					return nil, e
				}
			case enums.Remove:
				if e := db.Callback().Update().Remove(item.Name); e != nil {
					return nil, e
				}
			}
		}
	}

	if len(dbConfig.QueryCallbacks) > 0 {
		for _, item := range dbConfig.QueryCallbacks {
			switch item.OpType {
			case enums.Register:
				switch item.When {
				case enums.Before:
					if e := db.Callback().Query().Before(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				case enums.After:
					if e := db.Callback().Query().After(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				default:
					if e := db.Callback().Query().Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				}
			case enums.Replace:
				if e := db.Callback().Query().Replace(item.Name, item.Action); e != nil {
					return nil, e
				}
			case enums.Remove:
				if e := db.Callback().Query().Remove(item.Name); e != nil {
					return nil, e
				}
			}
		}
	}

	if len(dbConfig.DeleteCallbacks) > 0 {
		for _, item := range dbConfig.DeleteCallbacks {
			switch item.OpType {
			case enums.Register:
				switch item.When {
				case enums.Before:
					if e := db.Callback().Delete().Before(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				case enums.After:
					if e := db.Callback().Delete().After(item.Name).Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				default:
					if e := db.Callback().Delete().Register(item.RegisterName, item.Action); e != nil {
						return nil, e
					}
				}
			case enums.Replace:
				if e := db.Callback().Delete().Replace(item.Name, item.Action); e != nil {
					return nil, e
				}
			case enums.Remove:
				if e := db.Callback().Delete().Remove(item.Name); e != nil {
					return nil, e
				}
			}
		}
	}

	return db, nil
}

func getDialector(url string, dbType enums.DBType) (gorm.Dialector, error) {
	switch dbType {
	case enums.Sqlite:
		return nil, UnsupportedDBType

	case enums.Mysql:
		return mysql.Open(url), nil

	case enums.Postgresql:
		return postgres.Open(url), nil
	default:
		log.Debug("only support mysql, postgresql")
		return nil, UnsupportedDBType
	}

}
