package settings

type ApolloConfig struct {
	AppID         string `json:"appID" yaml:"appID"`                 // 应用ID
	Cluster       string `json:"cluster" yaml:"cluster"`             // 集群
	NamespaceName string `json:"namespaceName" yaml:"namespaceName"` // 命名空间
	IP            string `json:"ip" yaml:"ip"`                       // 服务器IP
	Secret        string `json:"secret" yaml:"secret"`               // 密钥
}
