package settings

// DB 数据库配置
// Mysql Url eg: root:my_pass@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=UTC
// Postgres Url eg: `host=127.0.0.1 user=root password=my_pass dbname=my_db port=5432 sslmode=disable TimeZone=UTC`
type DB struct {
	Url              string   `yaml:"url" json:"url"`                           // 写节点Url
	Replicas         []string `yaml:"replicas" json:"replicas"`                 // 读节点Url
	TablePrefix      string   `yaml:"tablePrefix" json:"tablePrefix"`           // 表前缀
	SingularTable    bool     `yaml:"singularTable" json:"singularTable"`       // 表复数禁用
	CreateBatchSize  int      `yaml:"createBatchSize" json:"createBatchSize"`   // 批量创建数量
	EnableRawSQL     bool     `yaml:"enableRawSql" json:"enableRawSql"`         // 打印原生SQL
	MaxFreeConnCount int      `yaml:"maxFreeConnCount" json:"maxFreeConnCount"` // 最大闲置连接数量
	MaxOpenConnCount int      `yaml:"maxOpenConnCount" json:"maxOpenConnCount"` // 最大连接数量
	ConnMaxLifetime  int      `yaml:"connMaxLifetime" json:"connMaxLifetime"`   // 连接存活最大时长, 单位: 秒
	FreeMaxLifetime  int      `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`   // 闲置连接存活的最大时间, 单位: 秒
}
