package utils

import (
	"ginLearn/structs/setting"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

var Log *logrus.Logger

func InitLog(conf setting.Config) {
	Log = logger(conf)
}

//logger logrus初始化设置
func logger(conf setting.Config) *logrus.Logger {
	var writers []io.Writer
	var f *os.File

	if IsNotEmptyStr(conf.Log.LogPath) {
		//确保日志目录存在
		dirLogPath := PathJoin(conf.Log.LogPath, conf.Application)
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

	if conf.Log.EnableLogConsole {
		log.Println("Log use console writer")
		writers = append(writers, os.Stdout)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	if len(writers) == 0 {
		log.Println("No set log writer, User console as default writer!!!")
		logger.Out = os.Stdout
	} else {
		logger.Out = io.MultiWriter(writers...)
	}

	gin.DefaultWriter = logger.Out
	gin.DefaultErrorWriter = logger.Out

	//设置日志级别
	level, err := logrus.ParseLevel(conf.Log.Level)

	if err != nil {
		logger.Fatalf("logger level convert failed!, err: %v", err)
	}

	logger.SetLevel(level)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}
