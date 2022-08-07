package server

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type HttpHandler func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

// GraceGrpc grpc启动配置.
type GraceGrpc struct {
	// Addr grpc地址, eg: ":7777"
	Addr   string
	Server *grpc.Server
	// RegisterReflect 启动grpc反射
	RegisterReflect bool
	// EnableGateway 启动grpc gateway
	EnableGateway bool
	// GatewayAddr gateway的地址
	GatewayAddr string
	// HttpHandlers http处理器
	HttpHandlers []HttpHandler
	// GatewayWaitTime gateway关闭时, 处理gateway请求时间
	GatewayWaitTime time.Duration
	// OnHookFunc 勾子函数
	OnHookFunc
}
