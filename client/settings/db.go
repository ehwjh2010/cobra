package settings

// DB 数据库配置
type DB struct {
	Host             string `yaml:"host" json:"host"`                         //数据库 IP
	Port             int    `yaml:"port" json:"port"`                         //数据库 端口
	User             string `yaml:"user" json:"user"`                         //用户名
	Password         string `yaml:"password" json:"password"`                 //密码
	Database         string `yaml:"database" json:"database"`                 //数据库名
	Location         string `yaml:"location" json:"location"`                 //时区
	TablePrefix      string `yaml:"tablePrefix" json:"tablePrefix"`           //表前缀
	SingularTable    bool   `yaml:"singularTable" json:"singularTable"`       //表复数禁用
	CreateBatchSize  int    `yaml:"createBatchSize" json:"createBatchSize"`   //批量创建数量
	EnableRawSQL     bool   `yaml:"enableRawSql" json:"enableRawSql"`         //打印原生SQL
	MaxFreeConnCount int    `yaml:"maxFreeConnCount" json:"maxFreeConnCount"` //最大闲置连接数量
	MaxOpenConnCount int    `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` //最大连接数量
	ConnMaxLifetime  int    `yaml:"connMaxLifetime" json:"connMaxLifetime"`   //连接存活最大时长, 单位: 秒
	FreeMaxLifetime  int    `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   //闲置连接存活的最大时间, 单位: 秒
}
