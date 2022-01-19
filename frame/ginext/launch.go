package ginext

import (
	"fmt"
	"github.com/ehwjh2010/viper"
	"github.com/ehwjh2010/viper/client/setting"
	"github.com/ehwjh2010/viper/frame/ginext/middleware"
	"github.com/ehwjh2010/viper/global"
	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/routine"
	"github.com/ehwjh2010/viper/server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	engine  *gin.Engine
	setting setting.Setting
}

func Viper(settings setting.Setting) *App {
	SetMode(settings.Debug)

	if err := log.InitLog(settings.LogConfig, settings.Application); err != nil {
		log.FatalErr("Log init failed", err)
	}

	if err := RegisterTrans(settings.Language); err != nil {
		log.FatalErr("Register validator translator failed, ", err)
	}

	newOnStartUp := make([]func() error, len(settings.OnStartUp)+1)

	newOnStartUp[0] = routine.SetUpDefaultTask(settings.Routine)

	copy(newOnStartUp[1:], settings.OnStartUp)

	settings.OnStartUp = newOnStartUp

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
	log.Infof(viper.SIGN + "\n" + "Viper Version: " + viper.VERSION)

	if app.setting.Swagger {
		log.Info("Use swagger, url: " +
			fmt.Sprintf(
				"http://%s:%d%s",
				app.setting.Host, app.setting.Port, global.SwaggerAPIUrl))
	}

	server.GraceServer(
		app.engine,
		app.setting.Host,
		app.setting.Port,
		app.setting.ShutDownTimeout,
		app.setting.OnStartUp,
		app.setting.OnShutDown)
}

//Engine 返回引擎
func (app *App) Engine() *gin.Engine {
	return app.engine
}
