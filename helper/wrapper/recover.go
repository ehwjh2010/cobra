package wrapper

import (
	"github.com/ehwjh2010/viper/log"
	"runtime/debug"
)

//type NoInputWithErr func() error
//
//type NoInputNoOutput func()

func PanicHandler() {
	if e := recover(); e != nil {
		log.Errorf("catch panic, %v\n%s", e, string(debug.Stack()))
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
