package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

const DefaultCreateBatchSize = 1000

type Mysql struct{}

func NewMysql() *Mysql {
	return &Mysql{}
}

func (m *Mysql) InitDB(dbConfig *DBConfig) (*gorm.DB, error) {
	dsn := dbConfig.Dsn()

	var sqlLogger = logger.Silent
	if dbConfig.EnableRawSQL {
		sqlLogger = logger.Info
	}

	var createBatchSize = DefaultCreateBatchSize
	if dbConfig.CreateBatchSize > 0 {
		createBatchSize = dbConfig.CreateBatchSize
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(dbConfig.MaxFreeConnCount)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConnCount)

	// SetConnMaxIdleTime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxIdleTime(dbConfig.FreeMaxLifetime * time.Minute)

	return db, nil
}

func (m *Mysql) Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	s, err := db.DB()
	if err != nil {
		log.Printf("Close conn; get db failed!, err: %v", err)
		return err
	}

	err = s.Close()

	if err != nil {
		log.Println("Close mysql failed!")
	} else {
		log.Println("Close mysql success!")
	}

	return err
}
