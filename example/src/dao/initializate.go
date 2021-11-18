package dao

import (
	"ginLearn/conf"
	"ginLearn/util"
	"ginLearn/util/structutils"
	"log"
)

var (
	DBClient    *util.DBClient
	CacheClient *util.RedisClient
)

//LoadDB 加载DB
func LoadDB(config *conf.DBConfig) {

	dbConfig := util.NewDBConfig()

	structutils.CopyProperties(config, dbConfig)

	client, err := util.InitDB(dbConfig)

	if err != nil {
		log.Panicf("Load mysql failed!, err: %v", err)
	}

	DBClient = client
}

//CloseDB 关闭DB
func CloseDB() error {
	return DBClient.Close()
}

//LoadCache 加载缓存
func LoadCache(config *conf.CacheConfig) {

	cacheConfig := util.NewCacheConfig()

	structutils.CopyProperties(config, cacheConfig)

	client, err := util.InitCache(cacheConfig)
	if err != nil {
		log.Panicf("Load redis failed!, err: %v\n", err)
	}

	CacheClient = client
}

//CloseCache 关闭缓存
func CloseCache() error {
	return CacheClient.Close()
}
