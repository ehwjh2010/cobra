package mysql

import (
	"github.com/ehwjh2010/viper/db/rdb"
	"github.com/ehwjh2010/viper/enums"
	"github.com/ehwjh2010/viper/log"
)

// SetUp 初始化Mysql
func SetUp(dbConfig rdb.DB) (client *rdb.DBClient, err error) {

	db, err := rdb.InitDBWithGorm(dbConfig, enums.Mysql)

	if err != nil {
		log.Debug("connect db failed")
		return nil, err
	}

	log.Debug("connect db success")

	client = rdb.NewDBClient(db, enums.Mysql, dbConfig)

	client.WatchHeartbeat()

	return client, nil
}
