package cobra

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/extend"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/gin-gonic/gin"
)

const sign = " .----------------.  .----------------.  .----------------.  .----------------.  .----------------.\n| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |\n| |     ______   | || |     ____     | || |   ______     | || |  _______     | || |      __      | |\n| |   .' ___  |  | || |   .'    `.   | || |  |_   _ \\    | || | |_   __ \\    | || |     /  \\     | |\n| |  / .'   \\_|  | || |  /  .--.  \\  | || |    | |_) |   | || |   | |__) |   | || |    / /\\ \\    | |\n| |  | |         | || |  | |    | |  | || |    |  __'.   | || |   |  __ /    | || |   / ____ \\   | |\n| |  \\ `.___.'\\  | || |  \\  `--'  /  | || |   _| |__) |  | || |  _| |  \\ \\_  | || | _/ /    \\ \\_ | |\n| |   `._____.'  | || |   `.____.'   | || |  |_______/   | || | |____| |___| | || ||____|  |____|| |\n| |              | || |              | || |              | || |              | || |              | |\n| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |\n '----------------'  '----------------'  '----------------'  '----------------'  '----------------'"
const version = "v1.0.7"

func Launch(application string, mode string, logConfig client.Log, middlewares []gin.HandlerFunc) *gin.Engine {
	if err := log.InitLog(&logConfig, application); err != nil {
		log.Fatalf("Init log failed! %v", err)
	}

	gin.SetMode(mode)

	log.Info(sign)

	engine := gin.New()

	middleware.UseMiddlewares(engine, middlewares...)

	return engine
}

func Run(engine *gin.Engine, serverConfig client.Server, onStartUp []func() error, onShutDown []func() error) {
	extend.GraceServer(engine, serverConfig, onStartUp, onShutDown)
}
