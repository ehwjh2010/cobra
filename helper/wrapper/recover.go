package wrapper

import (
	"github.com/ehwjh2010/viper/log"
	"go.uber.org/zap"
	"runtime/debug"
)

func PanicHandler() {
	if e := recover(); e != nil {
		log.Error("catch panic, ", zap.Any("err", e), zap.ByteString("stack", debug.Stack()))
	}
}
