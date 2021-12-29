package postgresql

import (
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/db/rdb"
)

//SetUp 初始化Mysql
func SetUp(dbConfig *client.DB) (client *rdb.DBClient, err error) {

	db, err := rdb.InitDBWithGorm(dbConfig, rdb.Postgresql)

	if err != nil {
		return nil, err
	}

	client = rdb.NewDBClient(db, rdb.Postgresql, *dbConfig)

	client.WatchHeartbeat()

	return client, nil
}
