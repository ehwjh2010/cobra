package ginext

import (
	"fmt"

	"github.com/ehwjh2010/viper/component/routine"
	"github.com/ehwjh2010/viper/constant"
	"github.com/ehwjh2010/viper/frame/ginext/middleware"
	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	engine  *gin.Engine
	setting server.Setting
}

func Viper(settings server.Setting) *App {
	SetMode(settings.Debug)

	if err := log.InitLog(settings.LogConfig, settings.Application); err != nil {
		log.FatalErr("Log init failed", err)
	}

	gin.DisableConsoleColor()
	writer := log.GetWriter()
	if writer != nil {
		gin.DefaultWriter = writer
	}

	if err := RegisterTrans(settings.Language); err != nil {
		log.FatalErr("Register validator translator failed, ", err)
	}

	if settings.EnableRtPool {
		newOnStartUp := append(settings.StartUp, routine.SetUpDefaultTask(settings.Routine))
		settings.StartUp = newOnStartUp

		newOnShutdown := append(settings.ShutDown, routine.CloseDefaultTask)
		settings.ShutDown = newOnShutdown
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

// Run 启动.
func (app *App) Run() {
	addr := fmt.Sprintf("%s:%d", app.setting.Host, app.setting.Port)

	if app.setting.Swagger {
		log.Debug("Use swagger, url: " + fmt.Sprintf("http://%s%s", addr, constant.SwaggerAPIUrl))
	}

	s := &server.GraceHttp{
		Engine:     app.engine,
		Addr:       addr,
		WaitSecond: app.setting.ShutDownTimeout,
		OnHookFunc: app.setting.OnHookFunc,
		GraceGrpc:  app.setting.GrpcConf,
	}

	log.FatalE(server.GraceHttpServer(s))
}

// Engine 返回引擎.
func (app *App) Engine() *gin.Engine {
	return app.engine
}
