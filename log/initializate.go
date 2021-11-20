package log

import (
	"github.com/ehwjh2010/cobra/client"
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

var sugaredLogger = zap.S()
var zLogger = zap.L()

// InitLog 初始化Logger
func InitLog(config *client.Log, application string) (err error) {
	if config == nil {
		return
	}

	var writeSyncer zapcore.WriteSyncer

	writeSyncer = getWriters(config, application)

	if strutils.IsNotEmptyStr(config.FileDir) {
		realLogDir := pathutils.PathJoin(config.FileDir, application)
		if err = pathutils.MakeDirs(realLogDir); err != nil {
			return
		}
	}

	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg := zap.New(core, zap.AddCallerSkip(1))
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
