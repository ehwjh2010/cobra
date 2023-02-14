package wrapper

import (
	"runtime/debug"

	"github.com/ehwjh2010/viper/log"
)

func PanicHandler() {
	if e := recover(); e != nil {
		log.Errorf("catch panic, err: %s, stack: %s", e, debug.Stack())
	}
}
