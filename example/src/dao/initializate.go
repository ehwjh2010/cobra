package dao

import (
	"ginLearn/client"
	"ginLearn/db/cache"
	"ginLearn/db/rdb"
	"ginLearn/db/rdb/mysql"
	"log"
)

var (
	DBClient    *rdb.DBClient
	CacheClient *cache.RedisClient
)

//LoadDB 加载DB
func LoadDB(config *client.DBConfig) {

	dbClient, err := mysql.InitMysql(config)

	if err != nil {
		log.Panicf("Load mysql failed!, err: %v", err)
	}

	DBClient = dbClient
}

//CloseDB 关闭DB
func CloseDB() error {
	return DBClient.Close()
}

//LoadCache 加载缓存
func LoadCache(config *client.CacheConfig) {

	cacheClient, err := cache.InitCache(config)
	if err != nil {
		log.Panicf("Load redis failed!, err: %v\n", err)
	}

	CacheClient = cacheClient
}

//CloseCache 关闭缓存
func CloseCache() error {
	return CacheClient.Close()
}
