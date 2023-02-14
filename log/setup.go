package log

import (
	"io"
	"os"

	"github.com/ehwjh2010/viper/constant"
	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/ehwjh2010/viper/helper/file"
	"github.com/ehwjh2010/viper/helper/path"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DefaultFilename      = "application.log"
	DefaultCaller        = 1
	DefaultTimeFieldName = "time"
)

type getWriterReq struct {
	log         *ZapLogReq
	logfilePath string
}

// InitZapLogger 初始化 zap Logger.
func InitZapLogger(config ZapLogReq, application string) (*zap.Logger, error) {
	var realLogFilePath string
	if str.IsNotEmpty(config.FileDir) {
		logFilePath, err := path.Relative2Abs(config.FileDir)
		if err != nil {
			return nil, err
		}
		realLogDir := path.JoinPath(logFilePath, application)
		if err = path.MkDirs(realLogDir); err != nil {
			return nil, err
		}

		if str.IsEmpty(config.FileName) {
			config.FileName = DefaultFilename
		}

		realLogFilePath = path.JoinPath(realLogDir, config.FileName)
	}

	writeSyncer, err := getWriters(getWriterReq{
		log:         &config,
		logfilePath: realLogFilePath,
	})
	if err != nil {
		return nil, err
	}

	encoder := getEncoder(config.TimeFieldName, config.TimeLayout)
	var l = new(zapcore.Level)
	err = l.UnmarshalText(str.Char2Bytes(config.Level))
	if err != nil {
		return nil, err
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	// 由于外部使用的都是包装∫后的方法, 需要加上AddCallerSkip(1),
	// zap.AddStacktrace(zapcore.WarnLevel) 这个函数的行为会一旦打印指定级别及以上的日志时, 自动打印堆栈
	// lg := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
	if config.Caller <= 0 {
		config.Caller = DefaultCaller
	}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(config.Caller))
	return logger, nil
}

func getEncoder(timeFieldName string, timeLayout string) zapcore.Encoder {
	if str.IsEmpty(timeFieldName) {
		timeFieldName = DefaultTimeFieldName
	}

	if str.IsEmpty(timeLayout) {
		timeLayout = constant.DefaultTimePattern
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeLayout)
	encoderConfig.TimeKey = timeFieldName
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriters(req getWriterReq) (zapcore.WriteSyncer, error) {
	var writers []io.Writer

	if req.log.EnableConsole {
		writers = append(writers, os.Stdout)
	}

	if str.IsNotEmpty(req.log.FileDir) {
		if req.log.Rotated {
			writer := &lumberjack.Logger{
				Filename:   req.logfilePath,
				MaxSize:    req.log.MaxSize,
				MaxBackups: req.log.MaxBackups,
				MaxAge:     req.log.MaxAge,
				LocalTime:  req.log.LocalTime,
				Compress:   req.log.Compress,
			}
			writers = append(writers, writer)
		} else {
			fileWriter, err := file.OpenFile(req.logfilePath)
			if err != nil {
				return nil, err
			}
			writers = append(writers, fileWriter)
		}
	}

	if writers == nil {
		Debugf("not set log output, Use stdout as log output")
		writers = append(writers, os.Stdout)
	}

	writer := io.MultiWriter(writers...)

	return zapcore.AddSync(writer), nil
}
