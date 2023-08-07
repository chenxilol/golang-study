package core

import (
	"producter-consumer/out"
	"sync"
)

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
	for i = 1; i <= taskNum; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}
	close(wo)
}

func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}
func Exec() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		producer(taskCh)
	}()
	go func() {
		defer wg.Done()
		consumer(taskCh)
	}()
	wg.Wait()
}
