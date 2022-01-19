package mysql

import (
	"github.com/ehwjh2010/viper/client/enum"
	"github.com/ehwjh2010/viper/client/setting"
	"github.com/ehwjh2010/viper/client/verror"
	"github.com/ehwjh2010/viper/db/rdb"
	"github.com/ehwjh2010/viper/log"
)

//SetUp 初始化Mysql
func SetUp(dbConfig *setting.DB) (client *rdb.DBClient, err error) {

	if dbConfig == nil {
		return nil, verror.InvalidConfig
	}

	db, err := rdb.InitDBWithGorm(dbConfig, enum.Mysql)

	if err != nil {
		log.Debug("Connect db failed")
		return nil, err
	}

	log.Debug("Connect db success")

	client = rdb.NewDBClient(db, enum.Mysql, dbConfig)

	client.WatchHeartbeat()

	return client, nil
}
