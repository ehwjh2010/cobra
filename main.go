package main

import (
	"fmt"
	"ginLearn/middleware"
	"ginLearn/src/configure"
	"ginLearn/src/route"
	"ginLearn/utils"
	"github.com/gin-gonic/gin"
)

func setUp() {
	//加载配置
	configure.LoadConfig()
	utils.InitLog(configure.Conf.Application, configure.Conf.LogConfig)
}

func main() {

	setUp()

	server := gin.New()

	route.BindRoute(server)

	middleware.UseMiddleWares(server)

	// TODO Custom Recovery behavior

	addr := fmt.Sprintf(":%d", configure.Conf.ServerPort)

	server.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
