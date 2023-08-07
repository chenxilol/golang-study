package core

import (
	"fmt"
	"producter-consumer/out"
	"sync"
)

type Task1 struct {
	ID int64
}

const Task1Num int64 = 10000

var TaskCh = make(chan Task1, 10)

func producer1(wr chan<- Task1) {
	var i int64
	for i = 1; i <= Task1Num; i++ {
		t := Task1{
			ID: i,
		}
		wr <- t
	}
	close(wr)
}
func (t *Task1) run1() {
	out.Println(t.ID)
}
func consumer1(ro <-chan Task1) {
	for t := range ro {
		if t.ID != 0 {
			t.run1()
		}

	}
}

func Test() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer1(TaskCh)
	}(wg)
	var i int64
	for i = 0; i < taskNum; i++ {
		if i%100 == 0 {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				consumer1(TaskCh)
			}(wg)
		}

	}

	wg.Wait()
	fmt.Println("完成")
}
