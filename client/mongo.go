package client

type Mongo struct {
	//Uri mongo Uri example:"mongodb://localhost:27017"
	Uri string `json:"uri" yaml:"uri"`
	//Database 数据库
	Database string `json:"database" yaml:"database"`
	//MaxConnectCount 最大连接数量
	MaxConnectCount uint64 `json:"maxConnectCount" yaml:"maxConnectCount"`
	//MinConnectCount 最小连接数
	MinConnectCount uint64 `json:"minConnectCount" yaml:"minConnectCount"`
	//FreeMaxLifetime 闲置连接最大存活时间, 单位: 分钟
	FreeMaxLifetime int `yaml:"freeMaxLifetime" json:"freeMaxLifetime"`
}
