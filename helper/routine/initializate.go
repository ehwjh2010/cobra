package routine

import (
	"github.com/ehwjh2010/viper/client"
	"github.com/panjf2000/ants/v2"
	"time"
)

type FinSignal struct{}

func NewFin() FinSignal {
	return FinSignal{}
}

type Task struct {
	rawConfig client.Routine
	p         *ants.Pool
}

func newTask(rawConfig client.Routine, p *ants.Pool) *Task {
	return &Task{rawConfig: rawConfig, p: p}
}

type TaskFunc func()

// SetUp 初始化协程池
func SetUp(conf *client.Routine) (*Task, error) {
	if conf.MaxWorkerCount <= 0 {
		conf.MaxWorkerCount = 5000
	}

	if conf.FreeMaxLifetime <= 0 {
		conf.FreeMaxLifetime = 3600
	}

	if conf.PanicHandler == nil {
		conf.PanicHandler = func(i interface{}) {

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

//AddTask 添加任务, 如果有设置
func (task *Task) AddTask(taskFunc TaskFunc, t time.Duration) error {
	//return task.p.Submit(taskFunc)

	return nil
}

//AddTaskWithTimeout 添加任务
func (task *Task) AddTaskWithTimeout(taskFunc TaskFunc, t time.Duration) error {
	//return task.p.Submit(taskFunc)

	return nil
}

//wrapperTaskFunc 包装任务函数
func (task *Task) wrapperTaskFunc(fn TaskFunc) func() {
	return func() {
		fn()
	}
}
