package utils

import (
	"ginLearn/client/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

const DefaultCreateBatchSize = 1000

var Conn *gorm.DB

func InitMySQL(mysqlConfig *setting.MysqlConfig) {
	dsn := mysqlConfig.Dsn()

	var sqlLogger = logger.Silent
	if mysqlConfig.EnableRawSQL {
		sqlLogger = logger.Info
	}

	var createBatchSize = DefaultCreateBatchSize
	if mysqlConfig.CreateBatchSize > 0 {
		createBatchSize = mysqlConfig.CreateBatchSize
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//打印SQL
		Logger: logger.Default.LogMode(sqlLogger),
		NamingStrategy: schema.NamingStrategy{
			//指定表前缀
			TablePrefix: mysqlConfig.TablePrefix,
			//表复数禁用
			SingularTable: mysqlConfig.SingularTable,
		},
		//批量操作 每批数量
		CreateBatchSize: createBatchSize,
	})

	if err != nil {
		log.Fatalf("Connect mysql failed! err: %v", err)
	}

	log.Println("Connect mysql success!")

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Access sqlDB failed! err: %v", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxFreeConnCount)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConnCount)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(mysqlConfig.ConnMaxLifetime * time.Minute)

	Conn = db
}

func CloseMySQL() error {
	s, err := Conn.DB()
	if err != nil {
		Log.Errorf("Close conn; get db failed!, err: %v", err)
	}

	err = s.Close()

	if err != nil {
		Log.Errorln("Close mysql failed!")
	} else {
		Log.Infoln("Close mysql success!")
	}

	return err
}
