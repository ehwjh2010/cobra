package routine

import (
	"github.com/ehwjh2010/viper/client"
	"github.com/panjf2000/ants/v2"
)

type Task struct {
	rawConfig client.Routine
	p         *ants.Pool
}

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
func (task Task) Reboot() {
	task.p.Reboot()
}

//Delay 添加任务, 如果有设置
func (task *Task) Delay(taskFunc TaskFunc) error {
	return task.p.Submit(taskFunc)
}
