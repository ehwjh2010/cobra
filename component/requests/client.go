package requests

import (
	"net/http"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ehwjh2010/viper/helper/types"
	"github.com/ehwjh2010/viper/log"
	"github.com/ehwjh2010/viper/verror"
	"github.com/go-resty/resty/v2"
	"github.com/levigross/grequests"
)

type HTTPClient struct {
	session  *grequests.Session
	maxTries types.NullInt
	logger   log.Logger
	client   *resty.Client
	timeout  time.Duration
}

type HOpt func(client *HTTPClient)
type InvokeMethod func(url string, ro *grequests.RequestOptions) (*grequests.Response, error)

func HWithMaxRetry(maxTries int) HOpt {
	return func(client *HTTPClient) {
		client.maxTries = types.NewInt(maxTries)
	}
}

func HWithLogger(logger log.Logger) HOpt {
	return func(client *HTTPClient) {
		client.logger = logger
	}
}

func HWithTimeout(timeout time.Duration) HOpt {
	return func(client *HTTPClient) {
		client.timeout = timeout
	}
}

func NewHTTPClient(hOPts ...HOpt) *HTTPClient {
	cli := &HTTPClient{
		client: resty.New(),
	}

	for _, fn := range hOPts {
		fn(cli)
	}

	return cli
}

// 默认超时时间为3秒, 重试次数为0.
var defaultHTTPClient = NewHTTPClient()

// method 请求.
func (api *HTTPClient) method(method string, url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	var (
		retryTimes int
		response   *HTTPResponse
		err        error
	)

	switch {
	case !api.maxTries.IsNull():
		retryTimes = api.maxTries.GetValue()
	default:
	}

	// 请求函数
	fn := func() error {
		invokeMethod := api.getDestMethod(method)
		resp, execErr := invokeMethod(url, req.toInternal())
		if execErr != nil {
			return execErr
		}

		response = NewResponse(resp)
		return nil
	}

	// 尝试次数加上第一次等于总次数
	retryCnt := retryTimes + 1

	err = retry.Do(fn, retry.Attempts(uint(retryCnt)))

	if err != nil {
		return nil, err
	}

	return response, nil
}

// getDestMethod 获取目标方法.
func (api *HTTPClient) getDestMethod(method string) InvokeMethod {
	switch method {
	case http.MethodGet:
		return api.session.Get
	case http.MethodPost:
		return api.session.Post
	case http.MethodPut:
		return api.session.Put
	case http.MethodPatch:
		return api.session.Patch
	case http.MethodDelete:
		return api.session.Delete
	case http.MethodHead:
		return api.session.Head
	case http.MethodOptions:
		return api.session.Options
	default:
		panic(verror.UnsupportedMethod)
	}
}

func (api *HTTPClient) do(method string, rOpts ...ROpt) {
	request := NewRequest(rOpts...)
	api.client.GetClient().Do(request.toInternal(api.logger))
}

// Get GET请求方法.
func (api *HTTPClient) Get(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	api.client.GetClient().Do()
	return
}

// Post Post请求方法.
func (api *HTTPClient) Post(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.method(http.MethodPost, url, rOpts...)
}

// Patch PATCH请求方法.
func (api *HTTPClient) Patch(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.method(http.MethodPatch, url, rOpts...)
}

// Put PUT请求方法.
func (api *HTTPClient) Put(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.method(http.MethodPut, url, rOpts...)
}

// Delete DELETE请求方法.
func (api *HTTPClient) Delete(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.method(http.MethodDelete, url, rOpts...)
}

// Head HEAD请求方法.
func (api *HTTPClient) Head(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.method(http.MethodHead, url, rOpts...)
}

// Options OPTIONS请求方法.
func (api *HTTPClient) Options(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.method(http.MethodOptions, url, rOpts...)
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
