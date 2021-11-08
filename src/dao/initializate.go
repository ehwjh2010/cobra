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

func LoadDB(config *conf.DBConfig) {

	dbConfig := utils.NewDBConfig()

	utils.CopyProperty(config, dbConfig)

	client, err := utils.InitDB(dbConfig)

	if err != nil {
		log.Panicf("Load mysql failed!, err: %v", err)
	}

	DBClient = client
}

func CloseDB() error {
	return DBClient.Close()
}

func LoadCache(config *conf.CacheConfig) {

	cacheConfig := utils.NewCacheConfig()

	utils.CopyProperty(config, cacheConfig)

	client, err := utils.InitCache(cacheConfig)
	if err != nil {
		log.Panicf("Load redis failed!, err: %v\n", err)
	}

	CacheClient = client
}

func CloseCache() error {
	return CacheClient.Close()
}
