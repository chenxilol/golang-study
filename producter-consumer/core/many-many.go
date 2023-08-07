package core

import (
	"fmt"
	"producter-consumer/out"
	"time"
)

type Task3 struct {
	ID int64
}

func (t *Task3) run() {
	out.Println(t.ID)
}

const TaskNum3 int64 = 100000
const Nums3 int64 = 100

var TaskCh3 = make(chan Task3, 10)

// 给生产者和消费者发送通知，告知你不要在无限循环了
var done = make(chan struct{})

func producer3(wo chan<- Task3, done chan struct{}) {
	var i int64
	for {
		if i >= TaskNum3 {
			i = 0
		}
		i++
		t := Task3{
			ID: i,
		}
		// 写到这里，会导致卡死，因为当关闭管道的时候，还在写入数据
		//wo <- t
		select {
		case wo <- t:
		case <-done:
			out.Println("生产者退出")
			return
		}
	}
}

func consumer3(ro <-chan Task3, done chan struct{}) {
	for {
		select {
		case t := <-ro:
			if t.ID != 0 {
				t.run()
			}
		case <-done:
			// 因为在关闭通道的时候可能还有数据，因为通道可能是有缓冲的，无缓冲的不会有值，所以最后要在读一下
			for t := range ro {
				if t.ID != 0 {
					t.run()
				}

			}
			out.Println("消费者退出")
			return
		}
	}
}

func Exec3() {
	go producer3(TaskCh3, done)
	go producer3(TaskCh3, done)
	go producer3(TaskCh3, done)
	go consumer3(TaskCh3, done)
	go consumer3(TaskCh3, done)
	go consumer3(TaskCh3, done)
	go consumer3(TaskCh3, done)
	time.Sleep(time.Second * 5)
	close(done)
	close(TaskCh3)
	time.Sleep(time.Second * 5)
	fmt.Println(len(TaskCh3))
}
