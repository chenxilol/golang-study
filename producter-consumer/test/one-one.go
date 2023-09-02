package test

import "producter-consumer/out"

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)

const taskNum int64 = 10000

func producer(wo chan<- Task) {
	var i int64
	for i = 0; i <= taskNum; i++ {
		task := Task{
			ID: i,
		}
		wo <- task

	}
	return
}
func com(ro <-chan Task) {
	for i := range ro {
		i.run()
	}
}

func Ex() {
	go producer(taskCh)
	go com(taskCh)
}
