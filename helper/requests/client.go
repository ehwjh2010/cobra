package requests

import (
	"github.com/ehwjh2010/viper/global"
	"github.com/ehwjh2010/viper/routine"
	"github.com/levigross/grequests"
	"time"
)

type HTTPClient struct {
	session *grequests.Session
}

func NewHTTPClient(req *HTTPRequest) *HTTPClient {
	cli := &HTTPClient{
		session: grequests.NewSession(req.toInternal()),
	}
	return cli
}

//CronClearIdle 定时清理闲置连接
func (api *HTTPClient) CronClearIdle(task *routine.Task, interval time.Duration) error {
	return task.AsyncDO(func() {
		for {
			<-time.After(interval)
			api.session.CloseIdleConnections()
		}
	})
}

var defaultHTTPClient = NewHTTPClient(NewRequest(RWithTimeout(global.ThreeSecDur)))

//Method 请求
func (api *HTTPClient) Method(method string, url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := grequests.Req(method, url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Get GET请求方法
func (api *HTTPClient) Get(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := api.session.Get(url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Post Post请求方法
func (api *HTTPClient) Post(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := api.session.Post(url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Patch PATCH请求方法
func (api *HTTPClient) Patch(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := api.session.Patch(url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Put PUT请求方法
func (api *HTTPClient) Put(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := api.session.Put(url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Delete DELETE请求方法
func (api *HTTPClient) Delete(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := api.session.Delete(url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Head HEAD请求方法
func (api *HTTPClient) Head(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := api.session.Head(url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Options OPTIONS请求方法
func (api *HTTPClient) Options(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := api.session.Options(url, req.toInternal())

	response := NewResponse(resp)

	return response, err
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
