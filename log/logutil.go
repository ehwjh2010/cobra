package log

import (
	"fmt"
	"log"

	"github.com/ehwjh2010/viper/enums"
)

func methodLogf(level string, msg string, args ...interface{}) {
	log.Printf("level: %s, message: %s", level, fmt.Sprintf(msg, args...))
}

func methodLog(level string, args ...interface{}) {
	log.Println(fmt.Sprintf("level: %s, ", level), args)
}

//-----------------------------------------------------------
//Debug日志相关方法

func Debug(args ...interface{}) {
	methodLog(enums.DEBUG, args...)
}

// Debugf 打印Debug级别日志.
func Debugf(msg string, args ...interface{}) {
	methodLogf(enums.DEBUG, msg, args...)
}

//-----------------------------------------------------------
//Info日志相关方法

// Info 打印info级别日志.
func Info(args ...interface{}) {
	methodLog(enums.INFO, args...)
}

// Infof 打印info级别日志.
func Infof(msg string, args ...interface{}) {
	methodLogf(enums.INFO, msg, args...)
}

//-----------------------------------------------------------
//Warn日志相关方法

// Warn 打印Warn级别日志.
func Warn(args ...interface{}) {
	methodLog(enums.WARN, args...)
}

// Warnf 打印Warn级别日志.
func Warnf(msg string, args ...interface{}) {
	methodLogf(enums.WARN, msg, args...)
}

//-----------------------------------------------------------
//Error日志相关方法

// Error 打印Error级别日志.
func Error(args ...interface{}) {
	methodLog(enums.ERROR, args...)
}

// Errorf 格式化打印Error级别日志.
func Errorf(msg string, args ...interface{}) {
	methodLogf(enums.ERROR, msg, args...)
}

// Err 打印Error级别日志, 以及打印err.
func Err(msg string, err error) {
	Errorf(fmt.Sprintf("%s, error: %s", msg, err))
}

// Errors 只打印错误.
func Errors(errs ...error) {
	Errorf(fmt.Sprintf("errors: %+v", errs))
}

//-----------------------------------------------------------
//Fatal日志相关方法

// Fatal 打印Fatal级别日志.
func Fatal(args ...interface{}) {
	methodLog(enums.FATAL, args...)
}

// Fatalf 格式化打印Fatal级别日志.
func Fatalf(msg string, args ...interface{}) {
	methodLogf(enums.FATAL, msg, args...)
}

// FatalErr 打印Fatal级别日志, 以及打印err.
func FatalErr(err error) {
	Fatalf(fmt.Sprintf("error: %s", err))
}

//-----------------------------------------------------------
//Panic日志相关方法

// Panic 打印Panic级别日志.
func Panic(args ...interface{}) {
	methodLog(enums.PANIC, args...)
}

// Panicf 打印Panic级别日志.
func Panicf(msg string, args ...interface{}) {
	methodLogf(enums.PANIC, msg, args...)
}
