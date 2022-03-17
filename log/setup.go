package log

import (
	"github.com/ehwjh2010/viper/client/settings"
	"github.com/ehwjh2010/viper/global"
	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/ehwjh2010/viper/helper/file"
	"github.com/ehwjh2010/viper/helper/path"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

const (
	DefaultFilename = "application.log"
	DefaultCaller   = 1
)

var logger = zap.L()
var sugaredLogger = zap.S()
var realLogFilePath string

// InitLog 初始化Logger
func InitLog(config settings.Log, application string) error {

	if str.IsNotEmpty(config.FileDir) {
		logFilePath, err := path.Relative2Abs(config.FileDir)
		if err != nil {
			return err
		}
		realLogDir := path.JoinPath(logFilePath, application)
		if err := path.MakeDirs(realLogDir); err != nil {
			return err
		}

		if str.IsEmpty(config.FileName) {
			config.FileName = DefaultFilename
		}

		realLogFilePath = path.JoinPath(realLogDir, config.FileName)
	}

	writeSyncer, err := getWriters(&config)
	if err != nil {
		return err
	}

	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText(str.Str2Bytes(config.Level))
	if err != nil {
		return err
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	//由于外部使用的都是包装∫后的方法, 需要加上AddCallerSkip(1),
	//zap.AddStacktrace(zapcore.WarnLevel) 这个函数的行为会一旦打印指定级别及以上的日志时, 自动打印堆栈
	//lg := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
	if config.Caller <= 0 {
		config.Caller = DefaultCaller
	}
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(config.Caller))
	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.S()调用即可
	sugaredLogger = zap.S()
	return nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(global.DefaultTimePattern)
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriters(conf *settings.Log) (zapcore.WriteSyncer, error) {
	var writers []io.Writer

	if conf.EnableConsole {
		writers = append(writers, os.Stdout)
	}

	if str.IsNotEmpty(conf.FileDir) {
		if conf.Rotated {
			writer := &lumberjack.Logger{
				Filename:   realLogFilePath,
				MaxSize:    conf.MaxSize,
				MaxBackups: conf.MaxBackups,
				MaxAge:     conf.MaxAge,
				LocalTime:  conf.LocalTime,
				Compress:   conf.Compress,
			}
			writers = append(writers, writer)
		} else {
			writer, err := file.OpenFile(realLogFilePath)
			if err != nil {
				return nil, err
			}
			writers = append(writers, writer)
		}
	}

	if writers == nil {
		Warn("No set log output, Use stdout as log output")
		writers = append(writers, os.Stdout)
	}

	w := io.MultiWriter(writers...)

	gin.DefaultWriter = w
	gin.DisableConsoleColor()

	return zapcore.AddSync(w), nil
}
