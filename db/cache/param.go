package cache

// Cache 缓存配置
type Cache struct {
	Network          string `yaml:"network" json:"network"`                   // 网络类型, tcp or unix，默认tcp
	Addr             string `yaml:"addr" json:"addr"`                         // Redis地址 eg: localhost:6379
	User             string `yaml:"user" json:"user"`                         // 用户, redis6.0开始支持
	Pwd              string `yaml:"pwd" json:"pwd"`                           // 密码
	MinFreeConnCount int    `yaml:"minFreeConnCount" json:"minFreeConnCount"` // 最小闲置连接数, 默认是3
	MaxOpenConnCount int    `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` // 最大连接数, 默认是10
	FreeMaxLifetime  int    `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   // 闲置连接存活的最大时间, -1表示无限制, 默认是10分钟, 单位: 秒
	ConnMaxLifetime  int    `yaml:"connMaxLifetime" json:"connMaxLifetime"`   // 连接存活最大时长, 默认是1小时, 单位: 秒
	Database         int    `yaml:"database" json:"database"`                 // 数据库
	ConnectTimeout   int    `yaml:"connectTimeout" json:"connectTimeout"`     // 建立连接超时时间, 默认是5秒, 单位: 秒
	ReadTimeout      int    `yaml:"readTimeout" json:"readTimeout"`           // 读取超时时间, -1表示无超时, 默认是3秒, 单位: 秒
	WriteTimeout     int    `yaml:"writeTimeout" json:"writeTimeout"`         // 写超时时间, 默认和ReadTimeout保持一致, 单位: 秒
	BusyWaitTimeOut  int    `yaml:"busyWaitTimeOut" json:"busyWaitTimeOut"`   // 当所有连接都处在繁忙状态时, 客户端等待可用连接的最大等待时长，默认为3秒
	MaxRetries       int    `yaml:"maxRetries" json:"maxRetries"`             // 最大尝试次数, -1表示不重试, 默认是3
}
