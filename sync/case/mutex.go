package _case

import (
	"fmt"
	"sync"
)

func MutexCase() {
	//singleRoutine()
	//many()
	Sync()
}

// 单协程操作
func singleRoutine() {
	mp := make(map[string]int)
	data := []string{
		"A", "B", "C", "D",
	}
	for i := 0; i < 20; i++ {
		for _, datum := range data {
			_, ok := mp[datum]
			if !ok {
				mp[datum] = 0
			}
			mp[datum] += 1
		}
	}
	fmt.Println(mp)
}

// 多协程线程不安全
func many() {
	mp := make(map[string]int)
	data := []string{
		"A", "B", "C", "D",
	}

	for i := 0; i < 20; i++ {

		go func() {
			for _, datum := range data {
				_, ok := mp[datum]
				if !ok {
					mp[datum] = 0
				}
				mp[datum] += 1
			}
		}()
	}
	fmt.Println(mp)
}

// Sync 互斥锁协程安全操作
func Sync() {
	type SyncMul struct {
		mp map[string]int
		sync.Mutex
	}
	m := SyncMul{
		mp:    make(map[string]int, 0),
		Mutex: sync.Mutex{},
	}
	data := []string{
		"A", "B", "C", "D",
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			for _, datum := range data {

				_, ok := m.mp[datum]
				if !ok {
					m.mp[datum] = 0
				}
				m.mp[datum] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(m.mp)
}

// 读写锁
type cache struct {
	data map[string]int
	sync.RWMutex
}

func newCache() *cache {
	return &cache{
		data:    make(map[string]int, 0),
		RWMutex: sync.RWMutex{},
	}
}
