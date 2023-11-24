package requests

import (
	"errors"
	"github.com/ehwjh2010/viper/enums"
	"time"

	"github.com/ehwjh2010/viper/helper/types"
	"github.com/ehwjh2010/viper/log"
	"github.com/go-resty/resty/v2"
)

var RequestError = errors.New("RequestError")

type HTTPClient struct {
	maxTries      types.NullInt
	logger        log.Logger
	client        *resty.Client
	timeout       time.Duration
	retryWaitTime time.Duration
}

type HOpt func(client *HTTPClient)

func HWithMaxRetry(maxTries int) HOpt {
	return func(client *HTTPClient) {
		client.maxTries = types.NewInt(maxTries)
	}
}

func HWithRetryWaitTime(t time.Duration) HOpt {
	return func(client *HTTPClient) {
		client.retryWaitTime = t
	}
}

func HWithLogger(logger log.Logger) HOpt {
	return func(client *HTTPClient) {
		client.client.SetLogger(logger)
		client.logger = logger
	}
}

func HWithTimeout(timeout time.Duration) HOpt {
	return func(client *HTTPClient) {
		client.timeout = timeout
	}
}

func NewHTTPClient(hOPts ...HOpt) *HTTPClient {
	logger := log.NewStdLogger()
	cli := &HTTPClient{
		client: resty.New().SetLogger(logger),
		logger: logger,
	}

	for _, fn := range hOPts {
		fn(cli)
	}

	return cli
}

// 默认超时时间为3秒, 重试次数为0.
var defaultHTTPClient = NewHTTPClient(
	HWithLogger(log.NewStdLogger()),
	HWithTimeout(enums.ThreeSecD))

type retryConfig struct {
	access        bool
	retries       int
	retryWaitTime time.Duration
}

func (api *HTTPClient) getRetryConfig(request *HTTPRequest) *retryConfig {
	if request.Retries > 0 || (!api.maxTries.IsNull() && api.maxTries.GetValue() > 0) {
		retries := api.getRetries(request)
		retryWaitTime := api.getRetryWaitTime(request)
		return &retryConfig{
			access:        true,
			retries:       retries,
			retryWaitTime: retryWaitTime,
		}
	}

	return &retryConfig{}
}

func (api *HTTPClient) accessRetry(request *HTTPRequest) bool {
	if request.Retries > 0 || (!api.maxTries.IsNull() && api.maxTries.GetValue() > 0) {
		return true
	}

	return false
}

func (api *HTTPClient) getRetries(request *HTTPRequest) int {
	if request.Retries > 0 {
		return request.Retries
	} else if !api.maxTries.IsNull() && api.maxTries.GetValue() > 0 {
		return api.maxTries.GetValue()
	}
	return 0
}

func (api *HTTPClient) getRetryWaitTime(request *HTTPRequest) time.Duration {
	if request.RetryWaitTime > 0 {
		return request.RetryWaitTime
	} else if api.retryWaitTime > 0 {
		return api.retryWaitTime
	}
	return 100 * time.Millisecond
}

func (api *HTTPClient) getTimeout(request *HTTPRequest) time.Duration {
	if request.Timeout > 0 {
		return request.Timeout
	} else if api.timeout > 0 {
		return api.timeout
	}

	return 0
}

func (api *HTTPClient) do(method, url string, rOpts ...ROpt) (*HTTPResponse, error) {
	request := NewRequest(rOpts...)
	client := api.client

	retryConfig := api.getRetryConfig(request)
	if retryConfig.access {
		client.SetRetryCount(retryConfig.retries).
			SetRetryWaitTime(retryConfig.retryWaitTime).
			AddRetryCondition(func(response *resty.Response, err error) bool {
				if err != nil || response.IsError() {
					return true
				}
				return false
			})
	}

	if timeout := api.getTimeout(request); timeout > 0 {
		client.SetTimeout(timeout)
	}

	r := client.R()
	request.setAttributes(r)
	response, err := r.Execute(method, url)
	if err != nil {
		return nil, err
	}

	temp := response.Error()
	if temp != nil {
		api.logger.Errorf("RequestError, err: %v", temp)
		return nil, RequestError
	}
	return NewResponse(response), nil
}

// Get GET请求方法.
func (api *HTTPClient) Get(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.do(resty.MethodGet, url, rOpts...)
}

// Post Post请求方法.
func (api *HTTPClient) Post(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.do(resty.MethodPost, url, rOpts...)
}

// Patch PATCH请求方法.
func (api *HTTPClient) Patch(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.do(resty.MethodPatch, url, rOpts...)
}

// Put PUT请求方法.
func (api *HTTPClient) Put(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.do(resty.MethodPut, url, rOpts...)
}

// Delete DELETE请求方法.
func (api *HTTPClient) Delete(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.do(resty.MethodDelete, url, rOpts...)
}

// Head HEAD请求方法.
func (api *HTTPClient) Head(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.do(resty.MethodHead, url, rOpts...)
}

// Options OPTIONS请求方法.
func (api *HTTPClient) Options(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.do(resty.MethodOptions, url, rOpts...)
}

func Get(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return defaultHTTPClient.Get(url, rOpts...)
}

func Post(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return defaultHTTPClient.Post(url, rOpts...)
}

func Patch(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return defaultHTTPClient.Patch(url, rOpts...)
}

func Put(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return defaultHTTPClient.Put(url, rOpts...)
}

func Delete(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return defaultHTTPClient.Delete(url, rOpts...)
}

func Head(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return defaultHTTPClient.Head(url, rOpts...)
}

func Options(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return defaultHTTPClient.Options(url, rOpts...)
}
