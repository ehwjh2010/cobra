package apollo

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/ehwjh2010/viper/helper/basic/boolean"
	"github.com/ehwjh2010/viper/helper/basic/double"
	"github.com/ehwjh2010/viper/helper/basic/integer"
	"github.com/ehwjh2010/viper/helper/basic/str"
	"github.com/ehwjh2010/viper/helper/serialize"
	"github.com/ehwjh2010/viper/verror"
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

// GetString
// 获取配置.
func (i *Client) GetString(key string) (string, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return "", err
	}

	value, err := str.Any2Char(v)
	return value, err
}

func (i *Client) GetStringSlice(key string) ([]string, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return nil, err
	}

	strings, ok := v.([]string)
	if !ok {
		return nil, verror.CastStrSliceErr
	}
	return strings, nil
}

// GetInt 获取配置.
func (i *Client) GetInt(key string) (int, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value, err := integer.Any2Int(v)
	return value, err
}

// GetInt32 获取配置.
func (i *Client) GetInt32(key string) (int32, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value, err := integer.Any2Int32(v)
	return value, err
}

// GetInt64 获取配置.
func (i *Client) GetInt64(key string) (int64, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value, err := integer.Any2Int64(v)
	return value, err
}

// GetBool 获取配置.
func (i *Client) GetBool(key string) (bool, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return false, err
	}

	value, err := boolean.Any2Bool(v)
	return value, err
}

// GetFloat64 获取配置.
func (i *Client) GetFloat64(key string) (float64, error) {
	v, err := i.cache.Get(key)
	if err != nil {
		return 0, err
	}

	value, err := double.Any2Double(v)
	return value, err
}

// GetJson 获取配置.
func (i *Client) GetJson(key string, v interface{}) error {
	data, err := i.cache.Get(key)
	if err != nil {
		return err
	}

	value, err := str.Any2Char(data)
	if err != nil {
		return err
	}
	if err = serialize.UnmarshalStr(value, v); err != nil {
		return err
	}

	return nil
}
