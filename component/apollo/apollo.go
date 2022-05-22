package apollo

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"

	"github.com/ehwjh2010/viper/helper/serialize"
)

type Client struct {
	cli agollo.Client

	cache agcache.CacheInterface

	rawConfig *ApolloConfig
}

func NewClient(cli agollo.Client, rawConfig *ApolloConfig) *Client {

	cache := cli.GetConfigCache(rawConfig.NamespaceName)

	return &Client{cli: cli, rawConfig: rawConfig, cache: cache}
}

// GetString 获取配置
func (i *Client) GetString(key string) (string, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return "", err
	}

	value := v.(string)
	return value, err
}

// GetInt32 获取配置
func (i *Client) GetInt32(key string) (int, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value := v.(int)
	return value, err
}

// GetInt 获取配置
func (i *Client) GetInt(key string) (int, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value := v.(int)
	return value, err
}

// GetInt64 获取配置
func (i *Client) GetInt64(key string) (int, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value := v.(int)
	return value, err
}

// GetBool 获取配置
func (i *Client) GetBool(key string) (bool, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return false, err
	}

	value := v.(bool)
	return value, nil
}

// GetFloat64 获取配置
func (i *Client) GetFloat64(key string) (float64, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value := v.(float64)
	return value, nil
}

// GetJson 获取配置
func (i *Client) GetJson(key string, v interface{}) error {
	data, err := i.cache.Get(key)
	if err != nil {
		return err
	}

	value := data.(string)
	if err = serialize.UnmarshalStr(value, v); err != nil {
		return err
	}

	return nil
}
