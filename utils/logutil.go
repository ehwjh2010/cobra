package utils

type LogConfig struct {
	LogPath          string `yaml:"logPath" json:"logPath"`                   //日志文件目录
	Level            string `yaml:"level" json:"level"`                       //日志级别
	EnableLogConsole bool   `yaml:"enableLogConsole" json:"enableLogConsole"` //是否输出打终端
	AccessMethodRow  bool   `yaml:"accessMethodRow" json:"accessMethodRow"`   //是否打印方法以及所在行数
}

type LogConfigOption func(*LogConfig)

const (
	LogPath          = ""
	Level            = "DEBUG"
	EnableLogConsole = false
	AccessMethodRow  = false
)

func NewLogConfig(args ...LogConfigOption) (logConfig *LogConfig) {
	logConfig = &LogConfig{
		LogPath:          LogPath,
		Level:            Level,
		EnableLogConsole: EnableLogConsole,
		AccessMethodRow:  AccessMethodRow,
	}

	for _, arg := range args {
		arg(logConfig)
	}

	return logConfig
}

func LogConfigWithLogPath(path string) LogConfigOption {
	return func(c *LogConfig) {
		c.LogPath = path
	}
}

func LogConfigWithLevel(level string) LogConfigOption {
	return func(config *LogConfig) {
		config.Level = level
	}
}

func LogConfigWithConsole(console bool) LogConfigOption {
	return func(config *LogConfig) {
		config.EnableLogConsole = console
	}
}

func LogConfigWithRow(enable bool) LogConfigOption {
	return func(config *LogConfig) {
		config.AccessMethodRow = enable
	}
}

type LogInterface interface {
	SetUp(application string) error

	Close() error
}
