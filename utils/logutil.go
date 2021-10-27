package utils

import (
	"ginLearn/client/setting"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

var Log *logrus.Logger

func InitLog(application string, logConfig *setting.LogConfig) {
	Log = logger(application, logConfig)
}

//logger logrus初始化设置
func logger(application string, logConfig *setting.LogConfig) *logrus.Logger {
	var writers []io.Writer
	var f *os.File

	if IsNotEmptyStr(logConfig.LogPath) {
		//确保日志目录存在
		dirLogPath := PathJoin(logConfig.LogPath, application)
		err := MakeDirs(dirLogPath)
		if err != nil {
			log.Fatalf("Access log dir failed! err: %v", err)
		}

		//确保日志文件存在
		logFilePath := PathJoin(dirLogPath, "application.log")
		f, err = OpenFileWithAppend(logFilePath)
		if err != nil {
			log.Fatalf("Access log file failed! err: %v", err)
		}
	}

	if f != nil {
		log.Println("Log use file writer")
		writers = append(writers, f)
	}

	if logConfig.EnableLogConsole {
		log.Println("Log use console writer")
		writers = append(writers, os.Stdout)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	if len(writers) == 0 {
		log.Println("No set log writer, User console as default writer!!!")
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(io.MultiWriter(writers...))
	}

	gin.DefaultWriter = logger.Out
	gin.DefaultErrorWriter = logger.Out

	//设置日志级别
	level, err := logrus.ParseLevel(logConfig.Level)

	if err != nil {
		logger.Fatalf("logger level convert failed!, err: %v", err)
	}

	logger.SetLevel(level)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	//添加打印日志所在文件以及行数, 比较影响性能, 是否使用自行决定
	if logConfig.AccessMethodRow {
		logger.SetReportCaller(true)
	}

	return logger
}

func Info(info ...interface{}) {
	Log.WithFields(logrus.Fields{}).Infoln(info...)
}

func Infof(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Infof(format, args...)
}
