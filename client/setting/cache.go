package setting

//Cache 缓存配置
type Cache struct {
	Host             string `yaml:"host" json:"host"`                         //Redis IP
	Port             int    `yaml:"port" json:"port"`                         //Redis 端口
	User             string `yaml:"user" json:"user"`                         //用户
	Pwd              string `yaml:"pwd" json:"pwd"`                           //密码
	MinFreeConnCount int    `yaml:"minFreeConnCount" json:"minFreeConnCount"` //最小闲置连接数
	MaxOpenConnCount int    `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数, 默认是每核10个连接
	FreeMaxLifetime  int    `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   //闲置连接存活的最大时间, 单位: 分钟
	Database         int    `yaml:"database" json:"database"`                 //数据库
	ConnectTimeout   int    `yaml:"connectTimeout" json:"connectTimeout"`     //连接Redis超时时间, 单位: 秒
	ReadTimeout      int    `yaml:"readTimeout" json:"readTimeout"`           //读取超时时间, 单位: 秒
	WriteTimeout     int    `yaml:"writeTimeout" json:"writeTimeout"`         //写超时时间, 单位: 秒
	DefaultTimeOut   int    `yaml:"defaultTimeOut" json:"defaultTimeOut"`     //默认缓存时间, 单位: 秒
	MaxRetries       int    `yaml:"maxRetries" json:"maxRetries"`             //最大尝试次数, 默认是3
	MaxConnAge       int    `yaml:"maxConnAge" json:"maxConnAge"`             //连接存活最大时长, 单位: 分钟
}
