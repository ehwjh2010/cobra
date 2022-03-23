package requests

import (
	"github.com/avast/retry-go"
	"github.com/ehwjh2010/viper/client/enums"
	"github.com/ehwjh2010/viper/component/routine"
	"github.com/ehwjh2010/viper/global"
	"github.com/ehwjh2010/viper/helper/types"
	"github.com/ehwjh2010/viper/log"
	"github.com/levigross/grequests"
	"net/http"
	"time"
)

type HTTPClient struct {
	session  *grequests.Session
	maxTries types.NullInt
}

type HOpt func(client *HTTPClient)

func HWithReq(r *HTTPRequest) HOpt {
	return func(client *HTTPClient) {
		client.session = grequests.NewSession(r.toInternal())
	}
}

func HWithRetryTimes(maxTries int) HOpt {
	return func(client *HTTPClient) {
		client.maxTries = types.NewInt(maxTries)
	}
}

func NewHTTPClient(hOPts ...HOpt) *HTTPClient {
	cli := &HTTPClient{}

	for _, fn := range hOPts {
		fn(cli)
	}

	return cli
}

// CronClearIdle 定时清理闲置连接
func (api *HTTPClient) CronClearIdle(task *routine.Task, interval time.Duration) error {
	var err error

	clearFn := func() {

		defer func() {
			if e := recover(); e != nil {
				log.Errorf("httpClient cronClearIdle occur err, err ==> ", e)
			}
		}()

		for {
			<-time.After(interval)
			api.session.CloseIdleConnections()
		}
	}

	if task != nil {
		err = task.AsyncDO(clearFn)
	} else {
		go clearFn()
	}
	return err
}

// 默认超时时间为3秒, 重试次数为0
var defaultHTTPClient = NewHTTPClient(
	HWithReq(NewRequest(RWithTimeout(enums.ThreeSecD), RWithUserAgent(global.UserAgent))),
)

// Method 请求
func (api *HTTPClient) Method(method string, url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	var (
		retryTimes int
		response   *HTTPResponse
		err        error
	)

	switch {
	case !req.RetryTimes.IsNull():
		retryTimes = req.RetryTimes.GetValue()
	case !api.maxTries.IsNull():
		retryTimes = api.maxTries.GetValue()
	}

	// 请求函数
	fn := func() error {
		resp, err := grequests.Req(method, url, req.toInternal())
		if err != nil {
			return err
		}

		response = NewResponse(resp)
		return nil
	}

	if retryTimes > 0 {
		err = retry.Do(fn, retry.Attempts(uint(retryTimes)))
	} else {
		err = fn()
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get GET请求方法
func (api *HTTPClient) Get(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodGet, url, rOpts...)
}

// Post Post请求方法
func (api *HTTPClient) Post(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodPost, url, rOpts...)
}

// Patch PATCH请求方法
func (api *HTTPClient) Patch(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodPatch, url, rOpts...)
}

// Put PUT请求方法
func (api *HTTPClient) Put(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodPut, url, rOpts...)
}

// Delete DELETE请求方法
func (api *HTTPClient) Delete(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodDelete, url, rOpts...)
}

// Head HEAD请求方法
func (api *HTTPClient) Head(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodHead, url, rOpts...)
}

// Options OPTIONS请求方法
func (api *HTTPClient) Options(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodOptions, url, rOpts...)
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
