package postgresql

import (
	"github.com/ehwjh2010/viper/client/enum"
	"github.com/ehwjh2010/viper/client/setting"
	"github.com/ehwjh2010/viper/client/verror"
	"github.com/ehwjh2010/viper/db/rdb"
)

//SetUp 初始化Mysql
func SetUp(dbConfig *setting.DB) (client *rdb.DBClient, err error) {

	if dbConfig == nil {
		return nil, verror.InvalidConfig
	}


	db, err := rdb.InitDBWithGorm(dbConfig, enum.Postgresql)

	if err != nil {
		return nil, err
	}

	client = rdb.NewDBClient(db, enum.Postgresql, dbConfig)

	client.WatchHeartbeat()

	return client, nil
}
