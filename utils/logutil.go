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

//InitLogrus logrus初始化设置
func InitLogrus(application string, logConfig *setting.LogConfig) {
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

	//设置gin框架相关日志输出
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

	Log = logger
}

//-----------------------------------------------------------

//Info日志相关方法

//Info 打印info级别日志
func Info(info ...interface{}) {
	Log.Infoln(info...)
}

//InfoWithFields 打印info级别日志, 包含fields
func InfoWithFields(fields logrus.Fields, info ...interface{}) {
	Log.WithFields(fields).Info(info...)
}

//InfoF 格式化打印info级别日志
func InfoF(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Infof(format, args...)
}

//InfoFWithFields 格式化打印info级别日志, 包含fields
func InfoFWithFields(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Infof(format, args...)
}

//-----------------------------------------------------------

//Warn日志相关方法

//Warn 打印Warn级别日志
func Warn(Warn ...interface{}) {
	Log.Warnln(Warn...)
}

//WarnWithFields 打印Warn级别日志, 包含fields
func WarnWithFields(fields logrus.Fields, Warn ...interface{}) {
	Log.WithFields(fields).Warn(Warn...)
}

//WarnF 格式化打印Warn级别日志
func WarnF(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Warnf(format, args...)
}

//WarnFWithFields 格式化打印Warn级别日志, 包含fields
func WarnFWithFields(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Warnf(format, args...)
}

//-----------------------------------------------------------

//Error日志相关方法

//Error 打印Error级别日志
func Error(Error ...interface{}) {
	Log.Errorln(Error...)
}

//ErrorWithFields 打印Error级别日志, 包含fields
func ErrorWithFields(fields logrus.Fields, Error ...interface{}) {
	Log.WithFields(fields).Error(Error...)
}

//ErrorF 格式化打印Error级别日志
func ErrorF(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Errorf(format, args...)
	Log.Debugln()
}

//ErrorFWithFields 格式化打印Error级别日志, 包含fields
func ErrorFWithFields(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Errorf(format, args...)
}

//-----------------------------------------------------------

//Fatal日志相关方法

//Fatal 打印Fatal级别日志
func Fatal(Fatal ...interface{}) {
	Log.Fatalln(Fatal...)
}

//FatalWithFields 打印Fatal级别日志, 包含fields
func FatalWithFields(fields logrus.Fields, Fatal ...interface{}) {
	Log.WithFields(fields).Fatal(Fatal...)
}

//FatalF 格式化打印Fatal级别日志
func FatalF(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Fatalf(format, args...)
	Log.Debugln()
}

//FatalFWithFields 格式化打印Fatal级别日志, 包含fields
func FatalFWithFields(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Fatalf(format, args...)
}
