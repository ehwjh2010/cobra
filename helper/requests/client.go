package requests

import (
	"github.com/levigross/grequests"
	"net/http"
)

type HTTPClient struct{}

//Method 请求
func (api *HTTPClient) Method(method string, url string, rOpts ...ROpt) (*HTTPResponse, error) {
	req := NewRequest(rOpts...)

	resp, err := grequests.Req(method, url, req.toInternal())

	response := NewResponse(resp)

	return response, err
}

//Get GET请求方法
func (api *HTTPClient) Get(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodGet, url, rOpts...)
}

//Post Post请求方法
func (api *HTTPClient) Post(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodPost, url, rOpts...)
}

//Patch PATCH请求方法
func (api *HTTPClient) Patch(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodPatch, url, rOpts...)
}

//Put PUT请求方法
func (api *HTTPClient) Put(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodPut, url, rOpts...)
}

//Delete DELETE请求方法
func (api *HTTPClient) Delete(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodDelete, url, rOpts...)
}

//Head HEAD请求方法
func (api *HTTPClient) Head(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodHead, url, rOpts...)
}

//Options OPTIONS请求方法
func (api *HTTPClient) Options(url string, rOpts ...ROpt) (*HTTPResponse, error) {
	return api.Method(http.MethodOptions, url, rOpts...)
}
