package wrapper

import (
	"runtime/debug"

	"github.com/ehwjh2010/viper/log"
	"go.uber.org/zap"
)

func PanicHandler() {
	if e := recover(); e != nil {
		log.Error("catch panic, ", zap.Any("err", e), zap.ByteString("stack", debug.Stack()))
	}
}
