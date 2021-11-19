package cobra

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/extend"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/gin-gonic/gin"
)

type App struct {
	Application string            `json:"application" yaml:"application"` //应用名
	Server      client.Server     `json:"server" yaml:"server"`           //server配置
	LogConfig   client.Log        `json:"logConfig" yaml:"logConfig"`     //日志配置
	OnStartUp   []func() error    //项目启动前执行函数
	OnShutDown  []func() error    //项目终止前执行函数
	Middlewares []gin.HandlerFunc //中间件
}

func (app *App) run() {

	if err := log.InitLog(&app.LogConfig, app.Application); err != nil {
		log.Panicf("Init log failed! %v", err)
	}

	engine := gin.New()

	middleware.UseMiddlewares(engine, app.Middlewares...)

	if app.Server.Port <= 0 {
		app.Server.Port = 8080
	}

	extend.GraceServer(engine, &app.Server, app.OnStartUp, app.OnShutDown)
}
