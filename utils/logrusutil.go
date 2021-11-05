package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

type Logrus struct {
	LogConfig    *LogConfig     `json:"logConfig"` // 日志配置
	file         *os.File       // 日志文件指针
	globalLogger *logrus.Logger // 全局logger
}

func NewLogrus(config *LogConfig) *Logrus {
	return &Logrus{LogConfig: config}
}

//SetUp logrus初始化设置
func (l *Logrus) SetUp(application string) error {
	var writers []io.Writer
	var f *os.File

	if IsNotEmptyStr(l.LogConfig.LogPath) {
		//确保日志目录存在
		dirLogPath := PathJoin(l.LogConfig.LogPath, application)
		err := MakeDirs(dirLogPath)
		if err != nil {
			//log.Fatalf("Access log dir failed! err: %v", err)
			return err
		}

		//确保日志文件存在
		logFilePath := PathJoin(dirLogPath, "application.log")
		f, err = OpenFileWithAppend(logFilePath)
		if err != nil {
			//log.Fatalf("Access log file failed! err: %v", err)
			return err
		}

	}

	if f != nil {
		log.Println("Log use file writer")
		writers = append(writers, f)
	}

	if l.LogConfig.EnableLogConsole {
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
	level, err := logrus.ParseLevel(l.LogConfig.Level)
	if err != nil {
		//logger.Fatalf("logger level convert failed!, err: %v", err)
		return err
	}
	logger.SetLevel(level)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	//添加打印日志所在文件以及行数, 比较影响性能, 是否使用自行决定
	if l.LogConfig.AccessMethodRow {
		logger.SetReportCaller(true)
	}

	l.globalLogger = logger
	l.file = f

	return nil
}

func (l *Logrus) Close() error {

	if l.file == nil {
		return nil
	}

	err := l.file.Close()

	if err == nil {
		log.Println("Close log file success")
	} else {
		log.Println("Close log file failed")
	}

	return err
}

//-----------------------------------------------------------

//Debug日志相关方法

//Debug 打印Debug级别日志
func (l *Logrus) Debug(Debug ...interface{}) {
	l.globalLogger.Debugln(Debug...)
}

//DebugWithFields 打印Debug级别日志, 包含fields
func (l *Logrus) DebugWithFields(fields logrus.Fields, Debug ...interface{}) {
	l.globalLogger.WithFields(fields).Debug(Debug...)
}

//DebugF 格式化打印Debug级别日志
func (l *Logrus) DebugF(format string, args ...interface{}) {
	l.globalLogger.Debugf(format, args...)
}

//DebugFWithFields 格式化打印Debug级别日志, 包含fields
func (l *Logrus) DebugFWithFields(format string, args ...interface{}) {
	l.globalLogger.WithFields(logrus.Fields{}).Debugf(format, args...)
}

//-----------------------------------------------------------

//Info日志相关方法

//Info 打印info级别日志
func (l *Logrus) Info(info ...interface{}) {
	l.globalLogger.Infoln(info...)
}

//InfoWithFields 打印info级别日志, 包含fields
func (l *Logrus) InfoWithFields(fields logrus.Fields, info ...interface{}) {
	l.globalLogger.WithFields(fields).Info(info...)
}

//InfoF 格式化打印info级别日志
func (l *Logrus) InfoF(format string, args ...interface{}) {
	l.globalLogger.Infof(format, args...)
}

//InfoFWithFields 格式化打印info级别日志, 包含fields
func (l *Logrus) InfoFWithFields(format string, args ...interface{}) {
	l.globalLogger.WithFields(logrus.Fields{}).Infof(format, args...)
}

//-----------------------------------------------------------

//Warn日志相关方法

//Warn 打印Warn级别日志
func (l *Logrus) Warn(Warn ...interface{}) {
	l.globalLogger.Warnln(Warn...)
}

//WarnWithFields 打印Warn级别日志, 包含fields
func (l *Logrus) WarnWithFields(fields logrus.Fields, Warn ...interface{}) {
	l.globalLogger.WithFields(fields).Warn(Warn...)
}

//WarnF 格式化打印Warn级别日志
func (l *Logrus) WarnF(format string, args ...interface{}) {
	l.globalLogger.Warnf(format, args...)
}

//WarnFWithFields 格式化打印Warn级别日志, 包含fields
func (l *Logrus) WarnFWithFields(format string, args ...interface{}) {
	l.globalLogger.WithFields(logrus.Fields{}).Warnf(format, args...)
}

//-----------------------------------------------------------

//Error日志相关方法

//Error 打印Error级别日志
func (l *Logrus) Error(Error ...interface{}) {
	l.globalLogger.Errorln(Error...)
}

//ErrorWithFields 打印Error级别日志, 包含fields
func (l *Logrus) ErrorWithFields(fields logrus.Fields, Error ...interface{}) {
	l.globalLogger.WithFields(fields).Error(Error...)
}

//ErrorF 格式化打印Error级别日志
func (l *Logrus) ErrorF(format string, args ...interface{}) {
	l.globalLogger.Errorf(format, args...)
}

//ErrorFWithFields 格式化打印Error级别日志, 包含fields
func (l *Logrus) ErrorFWithFields(format string, args ...interface{}) {
	l.globalLogger.WithFields(logrus.Fields{}).Errorf(format, args...)
}

//-----------------------------------------------------------

//Fatal日志相关方法

//Fatal 打印Fatal级别日志
func (l *Logrus) Fatal(Fatal ...interface{}) {
	l.globalLogger.Fatalln(Fatal...)
}

//FatalWithFields 打印Fatal级别日志, 包含fields
func (l *Logrus) FatalWithFields(fields logrus.Fields, Fatal ...interface{}) {
	l.globalLogger.WithFields(fields).Fatal(Fatal...)
}

//FatalF 格式化打印Fatal级别日志
func (l *Logrus) FatalF(format string, args ...interface{}) {
	l.globalLogger.Fatalf(format, args...)
}

//FatalFWithFields 格式化打印Fatal级别日志, 包含fields
func (l *Logrus) FatalFWithFields(format string, args ...interface{}) {
	l.globalLogger.WithFields(logrus.Fields{}).Fatalf(format, args...)
}

//-----------------------------------------------------------

//Panic日志相关方法

//Panic 打印Panic级别日志
func (l *Logrus) Panic(Panic ...interface{}) {
	l.globalLogger.Panicln(Panic...)
}

//PanicWithFields 打印Panic级别日志, 包含fields
func (l *Logrus) PanicWithFields(fields logrus.Fields, Panic ...interface{}) {
	l.globalLogger.WithFields(fields).Panic(Panic...)
}

//PanicF 格式化打印Panic级别日志
func (l *Logrus) PanicF(format string, args ...interface{}) {
	l.globalLogger.Panicf(format, args...)
}

//PanicFWithFields 格式化打印Panic级别日志, 包含fields
func (l *Logrus) PanicFWithFields(format string, args ...interface{}) {
	l.globalLogger.WithFields(logrus.Fields{}).Panicf(format, args...)
}
