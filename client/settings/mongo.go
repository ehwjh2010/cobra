package settings

type Mongo struct {
	//Uri mongo Uri example:"mongodb://localhost:27017"
	Uri string `json:"uri" yaml:"uri"`
	//Database 数据库
	Database string `json:"database" yaml:"database"`
	//MaxOpenConnCount 最大连接数量
	MaxOpenConnCount uint64 `json:"maxOpenConnCount" yaml:"maxOpenConnCount"`
	//MinOpenConnCount 最小连接数量
	MinOpenConnCount uint64 `json:"minOpenConnCount" yaml:"minOpenConnCount"`
	//FreeMaxLifetime 闲置连接最大存活时间, 单位: 秒
	FreeMaxLifetime int `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`
}
