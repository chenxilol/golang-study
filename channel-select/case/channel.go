package _case

import (
	"fmt"
	"time"
)

func Communication() {
	ch := make(chan int, 0)
	go Communication1(ch)
	go Communication2(ch)
}

//接受只写通道

func Communication1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

//接受只读通道

func Communication2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// 并发场景下的同步机制

func ConCurrentSync() {
	// 需要缓存的通道
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for c := range ch {
			fmt.Println(c)
		}
	}()
}

// NoticeAndMultiplexing 通知协程退出与多路复用
func NoticeAndMultiplexing() {
	ch := make(chan int, 0)
	strch := make(chan string, 0)
	done := make(chan struct{}, 0)
	go noticeAndMultiplexing1(ch)
	go noticeAndMultiplexing2(strch)
	// chan的多路复用
	go noticeAndMultiplexing3(ch, strch, done)
	time.Sleep(time.Second * 5)
	close(done)
}
func noticeAndMultiplexing1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
func noticeAndMultiplexing2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("数字: %d", i)
	}
}

// 多路复用
func noticeAndMultiplexing3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case str := <-strCh:
			fmt.Println(str)
		case <-done:
			fmt.Println("收到退出通知，退出当前协程")
			return
		}
		i++
		fmt.Println("累计执行次数: ", i)
	}
}
