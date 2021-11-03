package utils

import "github.com/sirupsen/logrus"

//-----------------------------------------------------------

//Debug日志相关方法

//Debug 打印Debug级别日志
func Debug(Debug ...interface{}) {
	Log.Debugln(Debug...)
}

//DebugWithFields 打印Debug级别日志, 包含fields
func DebugWithFields(fields logrus.Fields, Debug ...interface{}) {
	Log.WithFields(fields).Debug(Debug...)
}

//DebugF 格式化打印Debug级别日志
func DebugF(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Debugf(format, args...)
}

//DebugFWithFields 格式化打印Debug级别日志, 包含fields
func DebugFWithFields(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Debugf(format, args...)
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
}

//FatalFWithFields 格式化打印Fatal级别日志, 包含fields
func FatalFWithFields(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Fatalf(format, args...)
}

//-----------------------------------------------------------

//Panic日志相关方法

//Panic 打印Panic级别日志
func Panic(Panic ...interface{}) {
	Log.Panicln(Panic...)
}

//PanicWithFields 打印Panic级别日志, 包含fields
func PanicWithFields(fields logrus.Fields, Panic ...interface{}) {
	Log.WithFields(fields).Panic(Panic...)
}

//PanicF 格式化打印Panic级别日志
func PanicF(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Panicf(format, args...)
}

//PanicFWithFields 格式化打印Panic级别日志, 包含fields
func PanicFWithFields(format string, args ...interface{}) {
	Log.WithFields(logrus.Fields{}).Panicf(format, args...)
}
