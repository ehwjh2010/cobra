package server

import (
	"github.com/ehwjh2010/viper/client/verror"
)

type HookHandler func() error

// OnHookFunc 勾子结构体
type OnHookFunc struct {
	// StartUp 启动前执行函数
	StartUp []HookHandler
	// ShutDown 停止服务执行函数
	ShutDown []HookHandler
}

// invokeFunc 执行函数
func (h OnHookFunc) invokeFunc(functions []HookHandler) error {
	if functions == nil {
		return nil
	}

	var multiErr verror.MultiErr

	for _, function := range functions {
		if function == nil {
			continue
		}

		multiErr.AddErr(function())
	}

	return multiErr.AsStdErr()
}

// ExecuteStartUp 执行启动前操作
func (h OnHookFunc) ExecuteStartUp() error {
	return h.invokeFunc(h.StartUp)
}

// ExecuteShutDown 执行结束后操作
func (h OnHookFunc) ExecuteShutDown() error {
	return h.invokeFunc(h.ShutDown)
}
