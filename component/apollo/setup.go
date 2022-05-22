package apollo

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/component/log"
	"github.com/apolloconfig/agollo/v4/env/config"
)

// SetUp 初始化apollo
func SetUp(conf *ApolloConfig, logger log.LoggerInterface) (*Client, error) {

	if logger != nil {
		agollo.SetLogger(logger)
	}

	cli, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		c := &config.AppConfig{
			AppID:          conf.AppID,
			Cluster:        conf.Cluster,
			NamespaceName:  conf.NamespaceName,
			IP:             conf.IP,
			IsBackupConfig: false,
			Secret:         conf.Secret,
			MustStart:      true,
		}
		return c, nil
	})

	if err != nil {
		return nil, err
	}

	client := NewClient(cli, conf)

	return client, nil
}
