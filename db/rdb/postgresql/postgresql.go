package postgresql

import (
	"github.com/ehwjh2010/viper/client/enums"
	"github.com/ehwjh2010/viper/client/settings"
	"github.com/ehwjh2010/viper/db/rdb"
)

// SetUp 初始化Mysql
func SetUp(dbConfig settings.DB) (client *rdb.DBClient, err error) {
	db, err := rdb.InitDBWithGorm(dbConfig, enums.Postgresql)

	if err != nil {
		return nil, err
	}

	client = rdb.NewDBClient(db, enums.Postgresql, dbConfig)

	client.WatchHeartbeat()

	return client, nil
}
