package cobra

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/extend"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/gin-gonic/gin"
)

func Launch(application string, logConfig client.Log, middlewares []gin.HandlerFunc) *gin.Engine {

	if err := log.InitLog(&logConfig, application); err != nil {
		log.Fatalf("Init log failed! %v", err)
	}

	engine := gin.New()

	middleware.UseMiddlewares(engine, middlewares...)

	return engine
}

func Run(engine *gin.Engine, serverConfig client.Server, onStartUp []func() error, onShutDown []func() error) {
	extend.GraceServer(engine, serverConfig, onStartUp, onShutDown)
}
