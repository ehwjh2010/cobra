package main

import (
	"ginLearn/conf"
	"ginLearn/route"
	"ginLearn/utils"
)

func initialize() {
	conf.InitConfig()
	utils.InitLog(conf.Conf.Logfile)
}

func main() {
	initialize()

	r := route.GetRoute()

	// TODO add middleware
	// TODO Custom Recovery behavior

	r.Run("localhost:9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
