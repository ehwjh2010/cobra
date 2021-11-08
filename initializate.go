package main

import (
	"ginLearn/conf"
	"ginLearn/src/dao"
	"ginLearn/utils"
	"log"
)

func Initialization() {
	//保证项目resource初始化失败后, 资源会释放掉
	defer func() {
		if err := recover(); err != nil {
			releaseErrs := Release()
			if releaseErrs != nil {
				//log.Fatalf("Load resource failed!, %v, releaseErrs: %v\n", err, releaseErrs)
				log.Fatalf("Load resource failed!, %v, releaseErrs: %v\n", err, releaseErrs)
			} else {
				log.Fatalf("Load resource failed!, %v\n", err)
			}
		}
	}()

	conf.LoadConfig()

	LoadLog()

	dao.LoadDB(conf.Conf.DBConfig)

	dao.LoadCache(conf.Conf.CacheConfig)
}

//LoadLog 加载配置
func LoadLog() {

	logConfig := utils.NewLogConfig()

	err := utils.CopyProperty(conf.Conf.LogConfig, logConfig)
	if err != nil {
		log.Panicf(err.Error())
	}

	err = utils.InitLog(logConfig, conf.Conf.Application)

	if err != nil {
		log.Panicf("Load log failed!, err: %v", err)
	}

	log.Println("Load log success")
}

//Release 释放资源
func Release() (errs []error) {

	dbErr := dao.CloseDB()
	if dbErr != nil {
		errs = append(errs, dbErr)
	}

	cacheErr := dao.CloseCache()
	if cacheErr != nil {
		errs = append(errs, cacheErr)
	}

	return errs
}
