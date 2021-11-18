package mysql

import (
	"ginLearn/client"
	"ginLearn/db/rdb"
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
