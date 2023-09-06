package test

import (
	"fmt"
	"producter-consumer/out"
)

// Out 创建结构体，接受传来的chan
type Out struct {
	data chan interface{}
}

// 实例化Out
var outTest *Out

func NewOut() *Out {

	outTest = &Out{
		data: make(chan interface{}),
	}

	return outTest
}
func Print1(data interface{}) {
	outTest.data <- data
}
func (o *Out) Output() {
	for {
		select {
		case a := <-o.data:
			fmt.Println(a)
		}
	}
}

type Type struct {
	ID int
}

func (t *Type) run() {
	i := t.ID
	Print1(i)
}

var chanType = make(chan Type)
var done = make(chan struct{})

func Producer(wo chan<- Type, done <-chan struct{}) {
	for i := 0; i < 1000; i++ {
		t := Type{ID: i}
		select {
		case wo <- t:
		case <-done:
			fmt.Println("退出")
			return
		}
	}

}

func consumer3(ro <-chan Type, done chan struct{}) {
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
func Start() {
	go Producer(chanType, done)
	go Producer(chanType, done)
	go consumer3(chanType, done)
	go consumer3(chanType, done)
	go consumer3(chanType, done)
	go consumer3(chanType, done)
}
