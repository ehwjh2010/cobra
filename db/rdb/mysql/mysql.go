package mysql

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/db/rdb"
)

//InitMysql 初始化Mysql
func InitMysql(dbConfig *client.DB) (client *rdb.DBClient, err error) {

	db, err := rdb.InitDBWithGorm(dbConfig, rdb.Mysql)

	if err != nil {
		return nil, err
	}

	client = rdb.NewDBClient(db, rdb.Mysql)

	return client, nil
}
