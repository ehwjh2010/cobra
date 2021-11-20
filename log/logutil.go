package log

import "go.uber.org/zap"

//Debug日志相关方法

//Debug 打印Debug级别日志
func Debug(msg string, args ...zap.Field) {
	zap.L().Debug(msg, args...)
}

//Debugf 格式化打印Debug级别日志
func Debugf(format string, args ...interface{}) {
	zap.S().Debugf(format, args...)
}

//-----------------------------------------------------------

//Info日志相关方法

//Info 打印info级别日志
func Info(msg string, args ...zap.Field) {
	zap.L().Info(msg, args...)
}

//Infof 打印info级别日志
func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

//-----------------------------------------------------------

//Warn日志相关方法

//Warn 打印Warn级别日志
func Warn(msg string, args ...zap.Field) {
	zap.L().Warn(msg, args...)
}

//Warnf 格式化打印Warn级别日志
func Warnf(template string, args ...interface{}) {
	zap.S().Warnf(template, args...)
}

//-----------------------------------------------------------

//Error日志相关方法

//Error 打印Error级别日志
func Error(msg string, args ...zap.Field) {
	zap.L().Error(msg, args...)
}

//Errorf 格式化打印Error级别日志
func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args...)
}

//-----------------------------------------------------------

//Fatal日志相关方法

//Fatal 打印Fatal级别日志
func Fatal(msg string, args ...zap.Field) {
	zap.L().Fatal(msg, args...)
}

//Fatalf 格式化打印Fatal级别日志
func Fatalf(template string, args ...interface{}) {
	zap.S().Fatalf(template, args...)
}

//-----------------------------------------------------------

//Panic日志相关方法

//Panic 打印Panic级别日志
func Panic(msg string, args ...zap.Field) {
	zap.L().Panic(msg, args...)
}

//Panicf 格式化打印Panic级别日志
func Panicf(template string, args ...interface{}) {
	zap.S().Panicf(template, args...)
}
