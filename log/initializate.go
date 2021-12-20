package log

import (
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/global"
	"github.com/ehwjh2010/viper/util/file"
	"github.com/ehwjh2010/viper/util/path"
	"github.com/ehwjh2010/viper/util/str"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

const filename = "application.log"

var logger = zap.L()
var sugaredLogger = zap.S()

// InitLog 初始化Logger
func InitLog(config *client.Log, application string) (err error) {
	if config == nil {
		return
	}

	if str.IsNotEmpty(config.FileDir) {
		logFilePath, err := path.Relative2Abs(config.FileDir)
		if err != nil {
			return err
		}
		realLogDir := path.PathJoin(logFilePath, application)
		if err := path.MakeDirs(realLogDir); err != nil {
			return err
		}
	}

	var writeSyncer zapcore.WriteSyncer
	writeSyncer = getWriters(config, application)

	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	//由于外部使用的都是包装后的方法, 需要加上AddCallerSkip(1),
	//zap.AddStacktrace(zapcore.WarnLevel) 这个函数的行为会一旦打印指定级别及以上的日志时, 自动打印堆栈
	//lg := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.S()调用即可
	sugaredLogger = zap.S()
	return
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

func getWriters(conf *client.Log, application string) zapcore.WriteSyncer {
	var writers []io.Writer

	if conf.EnableConsole {
		writers = append(writers, os.Stdout)
	}

	if str.IsNotEmpty(conf.FileDir) {
		absPath, _ := path.Relative2Abs(conf.FileDir)
		filePath := path.PathJoin(absPath, application, filename)
		if conf.Rotated {
			writer := getRotedLogWriter(
				filePath,
				conf.MaxSize,
				conf.MaxBackups,
				conf.MaxAge,
				conf.LocalTime,
				conf.Compress)
			writers = append(writers, writer)
		} else {
			writer := getLogWriter(filePath)
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

	return zapcore.AddSync(w)
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
	f, _ := file.OpenFileWithAppend(filename)
	return f
}
