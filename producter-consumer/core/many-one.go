package core

import (
	"producter-consumer/out"
	"sync"
)

// 写线程当结束的时候会自动的结束协程，但是读不会，所以要在写完成后关闭流

type Task2 struct {
	ID int64
}

func (t *Task2) run() {
	out.Println(t.ID)
}

const TaskNum2 int64 = 1000
const Nums int64 = 100

var TaskCh2 = make(chan Task2, 10)

func producer2(wo chan<- Task2, startNum int64, nums int64) {
	var i int64
	for i = startNum; i < startNum+nums; i++ {
		t := Task2{
			ID: i,
		}
		wo <- t
	}

}
func consumer2(ro <-chan Task2) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}
func Exec2() {
	wg := &sync.WaitGroup{}
	pwg := &sync.WaitGroup{}
	var i int64
	for i = 0; i < TaskNum2; i += Nums {
		if i >= TaskNum2 {
			break
		}
		wg.Add(1)
		pwg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			defer pwg.Done()
			producer2(TaskCh2, i, Nums)

		}()
	}
	wg.Add(1)
	go func() {
		consumer2(TaskCh2)
		defer wg.Done()
	}()
	pwg.Wait()
	//加上go会避免因为时间差出现问题
	go close(TaskCh2)
	wg.Wait()
}
