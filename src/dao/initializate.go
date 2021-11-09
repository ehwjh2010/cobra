package dao

import (
	"ginLearn/conf"
	"ginLearn/utils"
	"log"
)

var (
	DBClient    *utils.DBClient
	CacheClient *utils.RedisClient
)

//LoadDB 加载DB
func LoadDB(config *conf.DBConfig) {

	dbConfig := utils.NewDBConfig()

	utils.CopyProperties(config, dbConfig)

	client, err := utils.InitDB(dbConfig)

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

	cacheConfig := utils.NewCacheConfig()

	utils.CopyProperties(config, cacheConfig)

	client, err := utils.InitCache(cacheConfig)
	if err != nil {
		log.Panicf("Load redis failed!, err: %v\n", err)
	}

	CacheClient = client
}

//CloseCache 关闭缓存
func CloseCache() error {
	return CacheClient.Close()
}
