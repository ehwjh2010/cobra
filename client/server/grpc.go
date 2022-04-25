package server

import (
	"google.golang.org/grpc"
)

// GraceGrpc grpc启动配置
type GraceGrpc struct {
	// Addr grpc地址, eg: ":7777"
	Addr   string
	Server *grpc.Server
	// RegisterReflect 启动grpc反射
	RegisterReflect bool
	// EnableGateway 启动grpc gateway, 该字段不会生效, 只是起到占位作用
	EnableGateway bool
	// GatewayAddr gateway的地址, eg: ":8888", 该字段不会生效, 只是起到占位作用
	GatewayAddr string
	OnHookFunc
}

func NewGraceGrpc() *GraceGrpc {
	return &GraceGrpc{}
}
