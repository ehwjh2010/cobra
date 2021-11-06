package resource

import (
	"log"
)

var releaseErrs []error

func init() {

	//保证项目resource初始化失败后, 资源会释放掉
	defer func() {
		if err := recover(); err != nil {
			Release()
			if releaseErrs != nil {
				//log.Fatalf("Load resource failed!, %v, releaseErrs: %v\n", err, releaseErrs)
				log.Fatalf("Load resource failed!, %v, releaseErrs: %v\n", err, releaseErrs)
			} else {
				log.Fatalf("Load resource failed!, %v\n", err)
			}
		}
	}()

	LoadConfig()

	LoadLog()

	LoadDB()

	LoadCache()
}

func Release() []error {
	mysqlErr := DBClient.Close()
	addErr(mysqlErr)

	redisErr := RedisClient.Close()
	addErr(redisErr)

	//...其他需要处理的数据

	return releaseErrs
}

func addErr(err error) {
	if err != nil {
		releaseErrs = append(releaseErrs, err)
	}
}
