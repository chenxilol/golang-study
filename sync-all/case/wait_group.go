package _case

import (
	"fmt"
	"sync"
	"time"
)

// StartNoGo 没有用并发
func StartNoGo() {
	a, b := 1000, 10000
	t := time.Now()
	for i := 0; i < 1000000000; i++ {
		multi(a, b)
	}
	totalTime := time.Since(t)
	fmt.Println(totalTime)
}

// StartGo 使用并发后,一定要习惯的使用sync.waitGroup ，最后使用wg.wait等待协程操作完成
// 使用并发的时候不要使用io流操作，例如打印等，因为这样会让原本并行的操作，变成串行
func StartGo() {
	a, b := 1000, 10000
	t := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 500000000; j++ {
				multi(a, b)
			}
		}()

	}
	wg.Wait()
	totalTime := time.Since(t)
	fmt.Println(totalTime)
}
func multi(a, b int) int {
	return a * b
}

func WaitGroupCase() {
	ch := make(chan []int, 1000)
	start := time.Now()
	wg := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		i := 0
		for item := range ch {
			fmt.Println(multi(item[0], item[1]))
			i++
		}
		time.Sleep(3 * time.Second)
		fmt.Println(i)
	}()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		wg2.Add(1)
		//保证了每个协程互不干涉
		// 用于协调多个 goroutine 的执行
		go func(wg1 *sync.WaitGroup) {
			defer wg1.Done()
			defer wg2.Done()
			for j := 0; j < 500; j++ {
				// 这里的i为什么等于2，
				//因为 i 是在一个循环中递增的，而 go func 可能在循环之后才运行。这可能导致 i 的值在 go func 执行时已经变为 2
				fmt.Println(i)
				ch <- []int{i, j}
			}
		}(wg)
	}
	wg.Wait()
	close(ch)
	wg2.Wait()
	t := time.Since(start)
	fmt.Println(t)
}
