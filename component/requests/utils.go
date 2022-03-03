package requests

import "net/http"

func CookieStr2Cookie(cookieStr string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}
	return request.Cookies()
}
