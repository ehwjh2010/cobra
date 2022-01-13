package routine

import (
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/log"
	"github.com/panjf2000/ants/v2"
	"time"
)

type TaskFunc func()

type AntsLogger func(string, ...interface{})

func (d AntsLogger) Printf(format string, args ...interface{}) {
	d(format, args...)
}

func defaultAntsLogger(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// SetUp 初始化协程池
func SetUp(conf *client.Routine) (*Task, error) {
	if conf.MaxWorkerCount <= 0 {
		conf.MaxWorkerCount = 10
	}

	if conf.FreeMaxLifetime <= 0 {
		conf.FreeMaxLifetime = 3600
	}

	if conf.Logger == nil {
		conf.Logger = AntsLogger(defaultAntsLogger)
	}

	if conf.PanicHandler == nil {
		conf.PanicHandler = func(i interface{}) {
			conf.Logger.Printf("execute task occur panic, panic ==> ", i)
		}
	}

	opts := ants.Options{
		ExpiryDuration:   time.Duration(conf.FreeMaxLifetime) * time.Second,
		PreAlloc:         true,
		MaxBlockingTasks: 0,
		Nonblocking:      false,
		Logger:           conf.Logger,
		PanicHandler:     conf.PanicHandler,
	}

	p, err := ants.NewPool(conf.MaxWorkerCount, ants.WithOptions(opts))
	if err != nil {
		return nil, err
	}

	task := newTask(*conf, p)

	return task, nil
}
