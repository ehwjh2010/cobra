package main

import (
	"fmt"
	"ginLearn/conf"
	"ginLearn/route"
	"ginLearn/utils"
)

func initialize() {
	conf.InitConfig()
	utils.InitLog(conf.Conf.Logfile, conf.Conf.EnableLogConsole)
}

func main() {
	initialize()

	r := route.GetRoute()

	utils.SetLogFormat(r)

	// TODO add middleware
	// TODO Custom Recovery behavior

	addr := fmt.Sprintf(":%d", conf.Conf.ServerPort)

	r.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
