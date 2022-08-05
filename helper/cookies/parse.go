package cookies

import "net/http"

type cookieParse struct{}

func NewCookieParser() CookieParser {
	return &cookieParse{}
}

func (c cookieParse) ParseRawCookie(rawCookie string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookie)
	request := http.Request{Header: header}
	cookies := request.Cookies()
	return cookies
}

func (c cookieParse) GetDestCookie(rawCookie, name string) (*http.Cookie, error) {
	header := http.Header{}
	header.Add("Cookie", rawCookie)
	request := http.Request{Header: header}
	value, err := request.Cookie(name)
	return value, err
}

func (c cookieParse) GetDestFromCookies(cookies []*http.Cookie, name string) (*http.Cookie, error) {
	request := http.Request{}
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	value, err := request.Cookie(name)
	return value, err
}

type CookieParser interface {
	ParseRawCookie(rawCookie string) []*http.Cookie
	GetDestCookie(rawCookie, name string) (*http.Cookie, error)
	GetDestFromCookies(cookies []*http.Cookie, name string) (*http.Cookie, error)
}
