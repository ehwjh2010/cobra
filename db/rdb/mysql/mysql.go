package mysql

import (
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/db/rdb"
	"github.com/ehwjh2010/viper/log"
)

//SetUp 初始化Mysql
func SetUp(dbConfig *client.DB) (client *rdb.DBClient, err error) {

	db, err := rdb.InitDBWithGorm(dbConfig, rdb.Mysql)

	if err != nil {
		log.Debug("Connect db failed")
		return nil, err
	}

	log.Debug("Connect db success")

	client = rdb.NewDBClient(db, rdb.Mysql, *dbConfig)

	client.WatchHeartbeat()

	return client, nil
}
