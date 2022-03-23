package requests

import (
	"context"
	"github.com/ehwjh2010/viper/helper/types"
	"github.com/levigross/grequests"
	"net/http"
	"time"
)

type HTTPRequest struct {
	Url        string            // Url
	Header     map[string]string // 请求头
	Params     map[string]string // 查询字符串
	Form       map[string]string // 表单参数
	Json       []byte            // json请求体
	Cookie     []*http.Cookie    // Cookie
	UserAgent  string            // UserAgent
	Timeout    time.Duration     // 请求超时时间
	Files      []*FileUpload     // 文件
	RetryTimes types.NullInt     // 重试次数
	Context    context.Context   // 上下文
}

func NewRequest(args ...ROpt) *HTTPRequest {
	req := &HTTPRequest{}
	for _, arg := range args {
		arg(req)
	}

	return req
}

type ROpt func(r *HTTPRequest)

func RWithContext(ctx context.Context) ROpt {
	return func(r *HTTPRequest) {
		r.Context = ctx
	}
}

// RWithRetry 设置重试次数, 必须>=0
func RWithRetry(times int) ROpt {
	return func(r *HTTPRequest) {
		if times < 0 {
			times = 0
		}
		r.RetryTimes = types.NewInt(times)
	}
}

func RWithHeader(header map[string]string) ROpt {
	return func(r *HTTPRequest) {
		r.Header = header
	}
}

func RWithParams(params map[string]string) ROpt {
	return func(r *HTTPRequest) {
		r.Params = params
	}
}

func RWithForm(form map[string]string) ROpt {
	return func(r *HTTPRequest) {
		r.Form = form
	}
}

func RWithJson(body []byte) ROpt {
	return func(r *HTTPRequest) {
		r.Json = body
	}
}

func RWithCookie(cookie []*http.Cookie) ROpt {
	return func(r *HTTPRequest) {
		r.Cookie = cookie
	}
}

func RWithUserAgent(userAgent string) ROpt {
	return func(r *HTTPRequest) {
		r.UserAgent = userAgent
	}
}

func RWithTimeout(timeout time.Duration) ROpt {
	return func(r *HTTPRequest) {
		r.Timeout = timeout
	}
}

func RWithFile(file *FileUpload) ROpt {
	return func(r *HTTPRequest) {
		r.Files = []*FileUpload{file}
	}
}

func RWithFiles(files []*FileUpload) ROpt {
	return func(r *HTTPRequest) {
		r.Files = files
	}
}

// toInternal 转换为RequestOptions
func (r *HTTPRequest) toInternal() *grequests.RequestOptions {
	if r == nil {
		return nil
	}

	rOpt := &grequests.RequestOptions{
		Headers:        r.Header,
		Params:         r.Params,
		Cookies:        r.Cookie,
		UserAgent:      r.UserAgent,
		RequestTimeout: r.Timeout,
		Context:        r.Context,
	}

	if r.Files != nil {
		rOpt.Files = BatchFileUpload2Internal(r.Files)
	}

	if r.Json != nil {
		rOpt.JSON = r.Json
	} else if r.Form != nil {
		rOpt.Data = r.Form
	}

	return rOpt
}
