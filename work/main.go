package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
	用map实现一个高并发

1. 面向高并发
2， 只存在插入操作和查询操作 O(1)
3. 查询时，若 key 存在，直接返回 val，若 key 不存在， 阻塞直到 key val 对 被放入后， 获取 val 返回；
*/
type MyConcurrentMap struct {
	MyMap map[int]string
	Rw    sync.RWMutex
}

func New() *MyConcurrentMap {
	return &MyConcurrentMap{
		MyMap: make(map[int]string),
		Rw:    sync.RWMutex{},
	}
}
func (m *MyConcurrentMap) Put(k int, v string) {
	m.Rw.Lock()
	defer m.Rw.Unlock()
	m.MyMap[k] = v
}

// , maxWaitingDuration time.Duration
func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (val string) {
	var value string
	m.Rw.RLock()
	defer m.Rw.RUnlock()
	wq := &sync.WaitGroup{}
	wq.Add(1)
	go func() {
		defer wq.Done()
		for {
			if val, ok := m.MyMap[k]; ok {
				value = val
				break
			}
		}
	}()
	wq.Wait()
	return value
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	b := []byte(s)
	newb := make([]byte, len(b))
	for i := len(b) - 1; i >= 0; i-- {
		newb[len(b)-i-1] = b[i]
	}
	return strings.EqualFold(s, string(newb[:]))
}

//func main() {
//	//	myMap := New()
//
//	//fmt.Println(isPalindrome(121))
//	fmt.Println(isPalindrome1(-121))
//	//	myMap.Put(1, "sadf")
//	//fmt.Println(myMap.Get(2, time.Second*4))
//}
//func isPalindrome1(x int) bool {
//	b := x % 11
//	if b == 0 {
//		return true
//	}
//	return false
//}

func main() {
	done := make(chan bool, 10)
	go func() {

		for {
			select {
			case <-done:
				fmt.Println("cancel")
				return
			default:
				fmt.Println("running")
				time.Sleep(time.Second * 3)
				fmt.Println("running6")
				done <- true
				fmt.Println("running7")
			}
		}
	}()

	time.Sleep(60 * time.Second)
	fmt.Println("main")
}
