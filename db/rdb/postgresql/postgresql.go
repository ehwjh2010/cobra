package postgresql

import (
	"github.com/ehwjh2010/viper/db/rdb"
	"github.com/ehwjh2010/viper/enums"
)

// SetUp 初始化Mysql
func SetUp(dbConfig rdb.DB) (client *rdb.DBClient, err error) {
	db, err := rdb.InitDBWithGorm(dbConfig, enums.Postgresql)

	if err != nil {
		return nil, err
	}

	client = rdb.NewDBClient(db, enums.Postgresql, dbConfig)

	client.WatchHeartbeat()

	return client, nil
}
