package wrapper

import (
	"github.com/ehwjh2010/viper/log"
	"go.uber.org/zap"
)

//type NoInputWithErr func() error
//
//type NoInputNoOutput func()

func PanicHandler() {
	if e := recover(); e != nil {
		log.Error("catch panic", zap.Any("panic", e))
	}
}

//func WrapRecoverWithError(fn NoInputWithErr) NoInputWithErr {
//	return func() error {
//		defer PanicHandler()
//		return fn()
//	}
//}
//
//func WrapRecoverWithSimple(fn NoInputNoOutput) NoInputNoOutput {
//	return func() {
//		defer PanicHandler()
//		fn()
//	}
//}
