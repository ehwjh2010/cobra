package requests

import (
	"context"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type HTTPRequest struct {
	Header    map[string]string // 请求头
	Params    map[string]string // 查询字符串
	Form      map[string]string // 表单参数
	Json      []byte            // json请求体
	Cookie    []*http.Cookie    // Cookie
	UserAgent string            // UserAgent
	Files     []*FileUpload     // 文件
	Context   context.Context   // 上下文
	Timeout   time.Duration     // 超時時間
	Retries   int               // 重試次數
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

func RWithTimeout(timeout time.Duration) ROpt {
	return func(r *HTTPRequest) {
		r.Timeout = timeout
	}
}

func RWithTries(retries int) ROpt {
	return func(r *HTTPRequest) {
		r.Retries = retries
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

// SetAttributes 設置屬性.
func (r *HTTPRequest) setAttributes(request *resty.Request) {
	if len(r.Header) > 0 {
		request.SetHeaders(r.Header)
	}

	if len(r.Params) > 0 {
		request.SetQueryParams(r.Params)
	}

	if len(r.Cookie) > 0 {
		request.SetCookies(r.Cookie)
	}

	if r.Context != nil {
		request.SetContext(r.Context)
	}

	if r.UserAgent != "" {
		request.SetHeader("UserAgent", r.UserAgent)
	}

	if len(r.Json) > 0 {
		request.SetBody(r.Json).
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json")
	}

	if len(r.Form) > 0 {
		request.SetFormData(r.Form)
	}

	if len(r.Files) > 0 {
		for _, f := range r.Files {
			request.SetMultipartField(f.FieldName, f.FileName, f.FileMime, f.FileContents)
		}
	}
}
