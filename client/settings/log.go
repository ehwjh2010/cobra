package settings

// Log 日志配置
type Log struct {
	Level         string `json:"level" yaml:"level"`                 // 日志级别
	EnableConsole bool   `yaml:"enableConsole" json:"enableConsole"` // 日志是否输出到终端
	Rotated       bool   `json:"rotated" yaml:"rotated"`             // 日志是否被分割
	FileDir       string `json:"fileDir" yaml:"fileDir"`             // 日志文件所在目录
	MaxSize       int    `json:"maxSize" yaml:"maxSize"`             // 每个日志文件长度的最大大小，默认100M
	MaxAge        int    `json:"maxAge" yaml:"maxAge"`               // 日志保留的最大天数(只保留最近多少天的日志)
	MaxBackups    int    `json:"maxBackups" yaml:"maxBackups"`       // 只保留最近多少个日志文件，用于控制程序总日志的大小
	LocalTime     bool   `json:"localtime" yaml:"localtime"`         // 是否使用本地时间，默认使用UTC时间
	Compress      bool   `json:"compress" yaml:"compress"`           // 是否压缩日志文件，压缩方法gzip
	Caller        int    `json:"caller" yaml:"caller"`               // 日志包装层数
	FileName      string `json:"fileName" yaml:"fileName"`           // 文件名
}
