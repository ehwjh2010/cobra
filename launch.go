package cobra

import (
	"fmt"
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/extend"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/ehwjh2010/cobra/util/validator"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

const SIGN = "\n .----------------.  .----------------.  .----------------.  .----------------.  .----------------.\n| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |\n| |     ______   | || |     ____     | || |   ______     | || |  _______     | || |      __      | |\n| |   .' ___  |  | || |   .'    `.   | || |  |_   _ \\    | || | |_   __ \\    | || |     /  \\     | |\n| |  / .'   \\_|  | || |  /  .--.  \\  | || |    | |_) |   | || |   | |__) |   | || |    / /\\ \\    | |\n| |  | |         | || |  | |    | |  | || |    |  __'.   | || |   |  __ /    | || |   / ____ \\   | |\n| |  \\ `.___.'\\  | || |  \\  `--'  /  | || |   _| |__) |  | || |  _| |  \\ \\_  | || | _/ /    \\ \\_ | |\n| |   `._____.'  | || |   `.____.'   | || |  |_______/   | || | |____| |___| | || ||____|  |____|| |\n| |              | || |              | || |              | || |              | || |              | |\n| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |\n '----------------'  '----------------'  '----------------'  '----------------'  '----------------'"
const VERSION = "v1.0.9"

type App struct {
	engine  *gin.Engine
	setting client.Setting
}

func Cobra(settings client.Setting) *App {
	SetMode(settings.Debug)

	if err := log.InitLog(&settings.LogConfig, settings.Application); err != nil {
		log.Fatal("Log init failed", zap.Error(err))
	}

	if err := validator.RegisterTrans(settings.Language); err != nil {
		log.Fatal("Register validator translator failed, ", zap.Error(err))
	}

	engine := gin.New()

	middleware.UseMiddlewares(engine, settings.Middlewares...)

	settings.Arrange()

	if settings.Swagger {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app := &App{
		engine:  engine,
		setting: settings,
	}
	return app
}

//Run 启动
func (app *App) Run() {
	log.Info(SIGN, zap.String("Version", "\n"+VERSION))

	if app.setting.Swagger {
		log.Info("Use swagger, url: " +
			fmt.Sprintf(
				"http://%s:%d%s",
				app.setting.Host, app.setting.Port, "/swagger/index.html"))
	}

	extend.GraceServer(
		app.engine,
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

//Engine 返回引擎
func (app *App) Engine() *gin.Engine {
	return app.engine
}
