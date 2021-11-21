package cobra

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/extend"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/gin-gonic/gin"
)

const SIGN = "\n .----------------.  .----------------.  .----------------.  .----------------.  .----------------.\n| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |\n| |     ______   | || |     ____     | || |   ______     | || |  _______     | || |      __      | |\n| |   .' ___  |  | || |   .'    `.   | || |  |_   _ \\    | || | |_   __ \\    | || |     /  \\     | |\n| |  / .'   \\_|  | || |  /  .--.  \\  | || |    | |_) |   | || |   | |__) |   | || |    / /\\ \\    | |\n| |  | |         | || |  | |    | |  | || |    |  __'.   | || |   |  __ /    | || |   / ____ \\   | |\n| |  \\ `.___.'\\  | || |  \\  `--'  /  | || |   _| |__) |  | || |  _| |  \\ \\_  | || | _/ /    \\ \\_ | |\n| |   `._____.'  | || |   `.____.'   | || |  |_______/   | || | |____| |___| | || ||____|  |____|| |\n| |              | || |              | || |              | || |              | || |              | |\n| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |\n '----------------'  '----------------'  '----------------'  '----------------'  '----------------'"
const VERSION = "v1.0.9"

type App struct {
	Engine *gin.Engine
	Application string
}


func Cobra(application string, debug bool, logConfig client.Log, middlewares []gin.HandlerFunc) *App {
	if err := log.InitLog(&logConfig, application); err != nil {
		log.Fatal(err.Error())
	}

	SetMode(debug)

	log.Info(SIGN)

	engine := gin.New()

	middleware.UseMiddlewares(engine, middlewares...)

	app := &App{
		Engine:      engine,
		Application: application,
	}

	return app
}

//Run 启动
func (app *App) Run(serverConfig client.Server, onStartUp func() error, onShutDown []func() error) {
	extend.GraceServer(app.Engine, serverConfig, onStartUp, onShutDown)
}

func SetMode(debug bool) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
