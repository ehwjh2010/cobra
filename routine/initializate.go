package routine

import (
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/log"
	"github.com/panjf2000/ants/v2"
	"time"
)

type Task struct {
	rawConfig client.Routine
	p         *ants.Pool
}

func newTask(rawConfig client.Routine, p *ants.Pool) *Task {
	return &Task{rawConfig: rawConfig, p: p}
}

type TaskFunc func()

type AntsLogger func(string, ...interface{})

func (d AntsLogger) Printf(format string, args ...interface{}) {
	d(format, args...)
}

func defaultAntsLogger(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// SetUp 初始化协程池
func SetUp(conf *client.Routine) (*Task, error) {
	if conf.MaxWorkerCount <= 0 {
		conf.MaxWorkerCount = 5000
	}

	if conf.FreeMaxLifetime <= 0 {
		conf.FreeMaxLifetime = 3600
	}

	if conf.Logger == nil {
		conf.Logger = AntsLogger(defaultAntsLogger)
	}

	if conf.PanicHandler == nil {
		conf.PanicHandler = func(i interface{}) {
			conf.Logger.Printf("err ==> ", i)
		}
	}

	opts := ants.Options{
		ExpiryDuration:   time.Duration(conf.FreeMaxLifetime) * time.Second,
		PreAlloc:         true,
		MaxBlockingTasks: 0,
		Nonblocking:      false,
		Logger:           conf.Logger,
	}

	p, err := ants.NewPool(conf.MaxWorkerCount, ants.WithOptions(opts))
	if err != nil {
		return nil, err
	}

	task := newTask(*conf, p)

	return task, nil
}

//Close 关闭协程池
func (task *Task) Close() {
	if task == nil || task.p == nil || task.p.IsClosed() {
		return
	}

	task.p.Release()
}

//Reboot 重启关闭的协程池
func (task Task) Reboot() {
	task.p.Reboot()
}

//Delay 添加任务, 如果有设置
func (task *Task) Delay(taskFunc TaskFunc) error {
	return task.p.Submit(taskFunc)
}
