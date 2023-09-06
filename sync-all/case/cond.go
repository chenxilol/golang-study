package _case

import (
	"fmt"
	"sync"
	"time"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})
	go readline(&list, cond)
	go readline(&list, cond)
	go readline(&list, cond)
	time.Sleep(3 * time.Second)
	initList(&list, cond)
}

// 这里为什么要加指针？
// 因为在golang中所有的传递都是值传递，引用传递只不过是复制的引用地址罢了，变量的地址已经改变，如果不用指针，在用append的时候会导致内存地址发生变化，会重新开辟一个地址
func initList(list *[]int, c *sync.Cond) {
	// 主叫方，可以持有锁，也可以不持有
	c.L.Lock()
	defer c.L.Unlock()
	for i := 0; i < 500; i++ {
		*list = append(*list, i)
	}
	// 唤醒所有线程
	c.Broadcast()
	// 唤醒其中一个线程
	//c.Signal()
}

// 被叫方必须加锁
func readline(list *[]int, c *sync.Cond) {
	c.L.Lock()
	defer c.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("readLine wait")
		c.Wait()
	}
	fmt.Println(*list)
}

// Queue  创建一个队列
type Queue struct {
	List []int
	Cond *sync.Cond
}

func NewQueue() *Queue {
	return &Queue{
		List: make([]int, 0),
		Cond: sync.NewCond(&sync.Mutex{}),
	}
}
func (q *Queue) Put(item int) {
	q.Cond.L.Lock()
	defer q.Cond.L.Unlock()
	q.List = append(q.List, item)
	// 通知被叫方释放协程
	q.Cond.Broadcast()
}
func (q *Queue) GetAny(n int) []int {
	q.Cond.L.Lock()
	defer q.Cond.L.Unlock()
	if len(q.List) < n {
		// 协程等待
		q.Cond.Wait()
	}
	List := q.List[:n]
	q.List = q.List[n:]

	return List
}
func CondQueueCase() {
	queue := NewQueue()
	// 创建多个协程读取队列
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 创建中间变量，因为在协程赋值的他的副本，如果不创建中间变量会导致数据竞争，以至于出现随机数
		i := i
		go func(wg1 *sync.WaitGroup) {
			defer wg1.Done()
			list := queue.GetAny(i)
			fmt.Printf("%d : %d \n", i, list)
		}(wg)

	}
	for i := 0; i < 100; i++ {
		queue.Put(i)
	}
	wg.Wait()
}
