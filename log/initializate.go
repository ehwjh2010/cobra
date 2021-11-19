package log

import (
	"github.com/ehwjh2010/cobra/util/fileutils"
	"github.com/ehwjh2010/cobra/util/pathutils"
	"github.com/ehwjh2010/cobra/util/strutils"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
)

const filename = "application.log"

var sugaredLogger *zap.SugaredLogger
var zLogger *zap.Logger

//LogConfig 日志配置
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

func NewLogConfig() *LogConfig {
	return &LogConfig{}
}

func (conf *LogConfig) RealLogDir(application string) string {
	return pathutils.PathJoin(conf.FileDir, application)
}

func (conf *LogConfig) FileName(application string) string {
	return pathutils.PathJoin(conf.FileDir, application, filename)
}

// InitLog 初始化Logger
func InitLog(config *LogConfig, application string) (err error) {
	var writeSyncer zapcore.WriteSyncer

	writeSyncer = getWriters(config, application)

	if strutils.IsNotEmptyStr(config.FileDir) {
		if err = pathutils.MakeDirs(config.RealLogDir(application)); err != nil {
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
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriters(conf *LogConfig, application string) zapcore.WriteSyncer {
	var writers []io.Writer

	if conf.EnableConsole {
		writers = append(writers, os.Stdout)
	}

	if strutils.IsNotEmptyStr(conf.FileDir) {
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

	if writers == nil {
		log.Println("No set log output, Use stdout as log output")
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
