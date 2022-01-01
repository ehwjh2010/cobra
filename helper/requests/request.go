package requests

import (
	"net/http"
	"time"
)

type Request struct {
	Url       string            //Url
	Header    map[string]string //请求头
	Params    map[string]string //查询字符串
	Form      map[string]string //表单参数
	Json      []byte            //json请求体
	Cookie    []*http.Cookie    //Cookie
	UserAgent string            //UserAgent
	Timeout   time.Duration     //请求超时时间
	Files     []FileUpload      //文件
}

func NewRequest(args ...ROpt) *Request {
	req := &Request{}
	for _, arg := range args {
		arg(req)
	}

	return req
}

type ROpt func(r *Request)

func RWithHeader(header map[string]string) ROpt {
	return func(r *Request) {
		r.Header = header
	}
}

func RWithParams(params map[string]string) ROpt {
	return func(r *Request) {
		r.Params = params
	}
}

func RWithForm(form map[string]string) ROpt {
	return func(r *Request) {
		r.Form = form
	}
}

func RWithJson(body []byte) ROpt {
	return func(r *Request) {
		r.Json = body
	}
}

func RWithCookie(cookie []*http.Cookie) ROpt {
	return func(r *Request) {
		r.Cookie = cookie
	}
}

func RWithUserAgent(userAgent string) ROpt {
	return func(r *Request) {
		r.UserAgent = userAgent
	}
}

func RWithTimeout(timeout time.Duration) ROpt {
	return func(r *Request) {
		r.Timeout = timeout
	}
}

func RWithFile(file FileUpload) ROpt {
	return func(r *Request) {
		r.Files = []FileUpload{file}
	}
}

func RWithFiles(files []FileUpload) ROpt {
	return func(r *Request) {
		r.Files = files
	}
}
