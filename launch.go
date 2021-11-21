package cobra

import (
	"fmt"
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/extend"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
)

const SIGN = "\n .----------------.  .----------------.  .----------------.  .----------------.  .----------------.\n| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |\n| |     ______   | || |     ____     | || |   ______     | || |  _______     | || |      __      | |\n| |   .' ___  |  | || |   .'    `.   | || |  |_   _ \\    | || | |_   __ \\    | || |     /  \\     | |\n| |  / .'   \\_|  | || |  /  .--.  \\  | || |    | |_) |   | || |   | |__) |   | || |    / /\\ \\    | |\n| |  | |         | || |  | |    | |  | || |    |  __'.   | || |   |  __ /    | || |   / ____ \\   | |\n| |  \\ `.___.'\\  | || |  \\  `--'  /  | || |   _| |__) |  | || |  _| |  \\ \\_  | || | _/ /    \\ \\_ | |\n| |   `._____.'  | || |   `.____.'   | || |  |_______/   | || | |____| |___| | || ||____|  |____|| |\n| |              | || |              | || |              | || |              | || |              | |\n| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |\n '----------------'  '----------------'  '----------------'  '----------------'  '----------------'"
const VERSION = "v1.0.9"

type App struct {
	Engine      *gin.Engine
	setting     client.Setting
}

func Cobra(settings client.Setting) *App {
	SetMode(settings.Debug)

	engine := gin.New()

	middleware.UseMiddlewares(engine, settings.Middlewares...)

	settings.Arrange()

	if settings.Swagger {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app := &App{
		Engine:  engine,
		setting: settings,
	}
	return app
}

//Run 启动
func (app *App) Run() {

	if err := log.InitLog(&app.setting.LogConfig, app.setting.Application); err != nil {
		log.Fatal("Log init failed", zap.Error(err))
	}

	log.Info(SIGN)

	log.Info("Use swagger, url: " +
		fmt.Sprintf(
			"http://%s:%d%s",
			app.setting.Host, app.setting.Port, "/swagger/index.html"))

	extend.GraceServer(
		app.Engine,
		app.setting.Host,
		app.setting.Port,
		app.setting.ShutDownTimeout,
		app.setting.OnStartUp,
		app.setting.OnShutDown)
}

func SetMode(debug bool) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
