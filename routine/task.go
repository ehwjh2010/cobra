package routine

import (
	"errors"
	"github.com/ehwjh2010/viper/client"
	"github.com/panjf2000/ants/v2"
	"sync"
)

type Task struct {
	rawConfig client.Routine
	p         *ants.Pool
}

var backgroundTask *Task
var mtx sync.Mutex

var NoEnableRoutinePool = errors.New("not enable routine pool, please add enableRtePool field to settings")

func newTask(rawConfig client.Routine, p *ants.Pool) *Task {
	return &Task{rawConfig: rawConfig, p: p}
}

//Close 关闭协程池
func (task *Task) Close() {
	if task == nil || task.p == nil || task.p.IsClosed() {
		return
	}

	task.p.Release()
}

//Reboot 重启关闭的协程池
func (task *Task) Reboot() {
	task.p.Reboot()
}

//AsyncDO 添加任务, 如果有设置
func (task *Task) AsyncDO(taskFunc TaskFunc) error {
	return task.p.Submit(taskFunc)
}

//CountInfo 获取协程池个数信息
func (task Task) CountInfo() map[string]int {
	result := make(map[string]int, 3)

	result["free"] = task.p.Free()
	result["running"] = task.p.Running()
	result["cap"] = task.p.Cap()

	return result
}

//SetUpDefaultTask 初始化后台任务
func SetUpDefaultTask(conf client.Routine) func() error {
	return func() error {
		if backgroundTask != nil {
			return nil
		}

		mtx.Lock()
		defer mtx.Unlock()
		if backgroundTask != nil {
			return nil
		}

		if task, err := SetUp(&conf); err != nil {
			return err
		} else {
			backgroundTask = task
		}
		return nil

	}
}

// AddTask 添加任务
func AddTask(taskFunc TaskFunc) error {
	if backgroundTask == nil {
		return NoEnableRoutinePool
	}

	return backgroundTask.AsyncDO(taskFunc)
}

//CountInfo 获取协程池个数信息
func CountInfo() (map[string]int, error) {
	if backgroundTask == nil {
		return nil, NoEnableRoutinePool
	}

	countInfo := backgroundTask.CountInfo()
	return countInfo, nil
}
