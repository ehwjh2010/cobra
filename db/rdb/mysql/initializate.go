package mysql

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/db/rdb"
)

//InitMysql 初始化Mysql
func InitMysql(dbConfig *client.DBConfig) (client *rdb.DBClient, err error) {

	db, err := InitMysqlWithGorm(dbConfig)

	if err != nil {
		return nil, err
	}

	client = rdb.NewDBClient(db, rdb.Mysql)

	return client, nil
}
