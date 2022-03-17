package routine

import (
	"github.com/ehwjh2010/viper/client/enums"
	"github.com/ehwjh2010/viper/client/settings"
	"github.com/ehwjh2010/viper/log"
	"github.com/panjf2000/ants/v2"
	"time"
)

type TaskFunc func()

func defaultAntsLogger(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// SetUp 初始化协程池
func SetUp(conf settings.Routine) (*Task, error) {

	if conf.MaxWorkerCount <= 0 {
		conf.MaxWorkerCount = 10
	}

	if conf.FreeMaxLifetime < 0 {
		conf.FreeMaxLifetime = enums.TwentyMinute
	}

	if conf.Logger == nil {
		conf.Logger = defaultAntsLogger
	}

	if conf.PanicHandler == nil {
		conf.PanicHandler = func(i interface{}) {
			conf.Logger.Printf("execute task occur panic, panic ==> ", i)
		}
	}

	opts := ants.Options{
		// 过期时间。表示 goroutine 空闲多长时间之后会被ants池回收
		ExpiryDuration: time.Duration(conf.FreeMaxLifetime) * time.Second,
		// 预分配。调用NewPool()/NewPoolWithFunc()之后预分配worker(管理一个工作 goroutine 的结构体)切片
		// 而且使用预分配与否会直接影响池中管理worker的结构
		PreAlloc: true,
		// 最大阻塞任务数量, 即池中 goroutine 数量已到池容量，且所有 goroutine 都处理繁忙状态，
		// 这时到来的任务会在阻塞列表等待。这个选项设置的是列表的最大长度。阻塞的任务数量达到这个值后，后续任务提交直接返回失败
		MaxBlockingTasks: 0,
		// 池是否阻塞，默认阻塞。提交任务时，如果ants池中 goroutine 已到上限且全部繁忙,
		// 阻塞的池会将任务添加的阻塞列表等待（当然受限于阻塞列表长度，见上一个选项）。非阻塞的池直接返回失败
		Nonblocking: false,
		// 日志记录器
		Logger: conf.Logger,
		// panic 处理. 遇到 panic 会调用这里设置的处理函数
		PanicHandler: conf.PanicHandler,
	}

	p, err := ants.NewPool(conf.MaxWorkerCount, ants.WithOptions(opts))
	if err != nil {
		return nil, err
	}

	task := newTask(conf, p)

	return task, nil
}
