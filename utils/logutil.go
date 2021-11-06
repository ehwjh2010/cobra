package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

const filename = "application.log"

var sugaredLogger *zap.SugaredLogger
var zLogger *zap.Logger

type LogConfig struct {
	//Level 日志级别
	Level string `json:"level" yaml:"level"`

	//EnableConsole 日志是否输出到终端
	EnableConsole bool `yaml:"enableConsole" json:"enableConsole"`

	//Rotated 日志是否被分割
	Rotated bool `json:"rotated" yaml:"rotated"`

	//FileDir 日志文件所在目录
	FileDir string `json:"fileDir" yaml:"fileDir"`

	//MaxSize 每个日志文件长度的最大大小，默认100M
	MaxSize int `json:"maxSize" yaml:"maxSize"`

	//MaxAge 日志保留的最大天数(只保留最近多少天的日志)
	MaxAge int `json:"maxAge" yaml:"maxAge"`

	//MaxBackups 只保留最近多少个日志文件，用于控制程序总日志的大小
	MaxBackups int `json:"maxBackups" yaml:"maxBackups"`

	//LocalTime 是否使用本地时间，默认使用UTC时间
	LocalTime bool `json:"localtime" yaml:"localtime"`

	//Compress 是否压缩日志文件，压缩方法gzip
	Compress bool `json:"compress" yaml:"compress"`
}

func (conf *LogConfig) RealLogDir(application string) string {
	return PathJoin(conf.FileDir, application)
}

func (conf *LogConfig) FileName(application string) string {
	return PathJoin(conf.FileDir, application, filename)
}

// InitLog 初始化Logger
func InitLog(config *LogConfig, application string) (err error) {
	var writeSyncer zapcore.WriteSyncer

	err = MakeDirs(config.RealLogDir(application))
	if err != nil {
		return
	}

	writeSyncer = getWriters(config, application)

	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.S()调用即可
	sugaredLogger = zap.S()
	zLogger = zap.L()
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriters(conf *LogConfig, application string) zapcore.WriteSyncer {
	var writers []io.Writer

	if conf.EnableConsole {
		writers = append(writers, os.Stdout)
	}

	if IsNotEmptyStr(conf.FileDir) {
		if conf.Rotated {
			writer := getRotedLogWriter(
				conf.FileName(application),
				conf.MaxSize,
				conf.MaxBackups,
				conf.MaxAge,
				conf.LocalTime,
				conf.Compress)
			writers = append(writers, writer)
		} else {
			writer := getLogWriter(conf.FileName(application))
			writers = append(writers, writer)
		}
	}

	return zapcore.AddSync(io.MultiWriter(writers...))
}

func getRotedLogWriter(filename string, maxSize, maxBackup, maxAge int, localTime bool, compress bool) io.Writer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		LocalTime:  localTime,
		Compress:   compress,
	}
	return lumberJackLogger
}

func getLogWriter(filename string) io.Writer {
	file, _ := OpenFileWithAppend(filename)
	return file
}

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

//PanicF 格式化打印Panic级别日志
func PanicF(format string, args ...interface{}) {
	sugaredLogger.Panicf(format, args...)
}
