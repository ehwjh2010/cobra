package server

import "net/http"

type GraceHttp struct {
	Engine http.Handler
	// Addr 地址
	Addr string `json:"addr" yaml:"addr"`
	// WaitSecond 等待时间
	WaitSecond int `json:"waitTime" yaml:"waitTime"`

	OnHookFunc
}
