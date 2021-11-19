package mysql

import (
	"cobra/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

const DefaultCreateBatchSize = 1000

func DefaultTimeFunc() time.Time {
	return time.Now().In(time.UTC)
}

func InitMysqlWithGorm(dbConfig *client.DBConfig) (*gorm.DB, error) {
	dsn := dbConfig.Dsn()

	var sqlLogger = logger.Silent
	if dbConfig.EnableRawSQL {
		sqlLogger = logger.Info
	}

	var createBatchSize = DefaultCreateBatchSize
	if dbConfig.CreateBatchSize > 0 {
		createBatchSize = dbConfig.CreateBatchSize
	}

	timeFunc := dbConfig.TimeFunc
	if timeFunc == nil {
		timeFunc = DefaultTimeFunc
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: timeFunc,
		//打印SQL
		Logger: logger.Default.LogMode(sqlLogger),
		NamingStrategy: schema.NamingStrategy{
			//指定表前缀
			TablePrefix: dbConfig.TablePrefix,
			//表复数禁用
			SingularTable: dbConfig.SingularTable,
		},
		//批量操作 每批数量
		CreateBatchSize: createBatchSize,
	})

	if err != nil {
		//log.Fatalf("Connect mysql failed! err: %v", err)
		return nil, err
	}

	log.Println("Connect mysql success!")

	sqlDB, err := db.DB()

	if err != nil {
		//log.Fatalf("Access sqlDB failed! err: %v", err)
		return nil, err
	}

	// SetMaxIdleConns 设置连接池中空闲连接最大数
	sqlDB.SetMaxIdleConns(dbConfig.MaxFreeConnCount)

	// SetMaxOpenConns 设置打开数据库最大连接数
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConnCount)

	// SetConnMaxIdleTime 设置闲置连接最长存活时间
	sqlDB.SetConnMaxIdleTime(dbConfig.FreeMaxLifetime * time.Minute)

	return db, nil
}
