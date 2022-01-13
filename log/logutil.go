package log

import "go.uber.org/zap"

//Debug日志相关方法

//Debug 打印Debug级别日志
func Debug(msg string, args ...zap.Field) {
	logger.Debug(msg, args...)
}

//Debugf 格式化打印Debug级别日志
func Debugf(format string, args ...interface{}) {
	sugaredLogger.Debugf(format, args...)
}

//-----------------------------------------------------------

//Info日志相关方法

//Info 打印info级别日志
func Info(msg string, args ...zap.Field) {
	logger.Info(msg, args...)
}

//Infof 打印info级别日志
func Infof(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

//-----------------------------------------------------------

//Warn日志相关方法

//Warn 打印Warn级别日志
func Warn(msg string, args ...zap.Field) {
	logger.Warn(msg, args...)
}

//Warnf 格式化打印Warn级别日志
func Warnf(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

//-----------------------------------------------------------

//Error日志相关方法

//Error 打印Error级别日志
func Error(msg string, args ...zap.Field) {
	logger.Error(msg, args...)
}

//Errorf 格式化打印Error级别日志
func Errorf(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

//Err 打印Error级别日志, 以及打印err
func Err(msg string, errs ...error) {
	fields := make([]zap.Field, len(errs))
	for index, err := range errs {
		fields[index] = zap.Error(err)
	}
	logger.Error(msg, fields...)
}

//-----------------------------------------------------------

//Fatal日志相关方法

//Fatal 打印Fatal级别日志
func Fatal(msg string, args ...zap.Field) {
	logger.Fatal(msg, args...)
}

//Fatalf 格式化打印Fatal级别日志
func Fatalf(template string, args ...interface{}) {
	sugaredLogger.Fatalf(template, args...)
}

//FatalErr 打印Fatal级别日志, 以及打印err
func FatalErr(msg string, errs ...error) {
	fields := make([]zap.Field, len(errs))
	for index, err := range errs {
		fields[index] = zap.Error(err)
	}
	logger.Fatal(msg, fields...)
}

//-----------------------------------------------------------

//Panic日志相关方法

//Panic 打印Panic级别日志
func Panic(msg string, args ...zap.Field) {
	logger.Panic(msg, args...)
}

//Panicf 格式化打印Panic级别日志
func Panicf(template string, args ...interface{}) {
	sugaredLogger.Panicf(template, args...)
}
