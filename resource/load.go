package resource

import "log"

func init() {

	//保证项目resource初始化失败后, 资源会释放掉
	defer func() {
		if err := recover(); err != nil {
			Release()
			log.Fatalf("Load resource failed!, %v\n", err)
		}
	}()

	LoadConfig()

	LoadLogrus()

	LoadMySQL()
}
