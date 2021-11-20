package log

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/enum"
	"github.com/ehwjh2010/cobra/util/fileutils"
	"github.com/ehwjh2010/cobra/util/pathutils"
	"github.com/ehwjh2010/cobra/util/strutils"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

const filename = "application.log"

// InitLog 初始化Logger
func InitLog(config *client.Log, application string) (err error) {
	if config == nil {
		return
	}

	if strutils.IsNotEmptyStr(config.FileDir) {
		realLogDir := pathutils.PathJoin(config.FileDir, application)
		if err = pathutils.MakeDirs(realLogDir); err != nil {
			return
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
	lg := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.S()调用即可
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(enum.DefaultTimePattern)
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

	if strutils.IsNotEmptyStr(conf.FileDir) {
		filePath := pathutils.PathJoin(conf.FileDir, application, filename)
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
	file, _ := fileutils.OpenFileWithAppend(filename)
	return file
}
