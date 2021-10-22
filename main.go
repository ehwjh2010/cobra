package main

import (
	"fmt"
	"ginLearn/conf"
	"ginLearn/middleware"
	"ginLearn/route"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
)

func initialize() {
	conf.InitConfig()
	utils.InitLog(
		conf.Conf.Application,
		conf.Conf.Log.LogPath,
		conf.Conf.Log.Level,
		conf.Conf.Log.EnableLogConsole,
	)
}

func main() {
	r := gin.New()

	initialize()

	route.BindRoute(r)

	middleware.UseMiddleWares(r)

	// TODO Custom Recovery behavior

	addr := fmt.Sprintf(":%d", conf.Conf.ServerPort)

	r.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
