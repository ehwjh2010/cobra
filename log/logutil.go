package log

import "go.uber.org/zap"

//Debug日志相关方法

//Debug 打印Debug级别日志
func Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

//Debugf 格式化打印Debug级别日志
func Debugf(format string, args ...interface{}) {
	sugaredLogger.Debugf(format, args...)
}

//-----------------------------------------------------------

//Info日志相关方法

//Info 打印info级别日志
func Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

//Infol 打印info级别日志
func Infol(msg string, fields ...zap.Field) {
	zLogger.Info(msg, fields...)
}

//Infof 格式化打印info级别日志
func Infof(format string, args ...interface{}) {
	sugaredLogger.Infof(format, args...)
}

//-----------------------------------------------------------

//Warn日志相关方法

//Warn 打印Warn级别日志
func Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

//Warnf 格式化打印Warn级别日志
func Warnf(format string, args ...interface{}) {
	sugaredLogger.Warnf(format, args...)
}

//-----------------------------------------------------------

//Error日志相关方法

//Error 打印Error级别日志
func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

//Errorl 打印Error级别日志
func Errorl(msg string, fields ...zap.Field) {
	zLogger.Error(msg, fields...)
}

//Errorf 格式化打印Error级别日志
func Errorf(format string, args ...interface{}) {
	sugaredLogger.Errorf(format, args...)
}

//-----------------------------------------------------------

//Fatal日志相关方法

//Fatal 打印Fatal级别日志
func Fatal(args ...interface{}) {
	sugaredLogger.Fatal(args...)
}

//Fatalf 格式化打印Fatal级别日志
func Fatalf(format string, args ...interface{}) {
	sugaredLogger.Fatalf(format, args...)
}

//-----------------------------------------------------------

//Panic日志相关方法

//Panic 打印Panic级别日志
func Panic(args ...interface{}) {
	sugaredLogger.Panic(args...)
}

//Panicf 格式化打印Panic级别日志
func Panicf(format string, args ...interface{}) {
	sugaredLogger.Panicf(format, args...)
}
