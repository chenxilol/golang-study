package main

import (
	"context"
	"example.com/myproject/case"
	"os"
	"os/signal"
)

func main() {
	// 读写锁，以及互斥锁
	//_case.MutexCase()
	// sync-all.map 适用于读多写少，不适用于计数器

	//_case.MapCase()
	//没有使用并发
	//_case.StartNoGo()
	// 使用并发后
	//_case.StartGo()
	// 使用waitGroup,对协程进行加锁，以及等待
	//  _case.WaitGroupCase()
	// 利用sync.cond 对协程进行选择性的唤醒
	//	_case.CondCase()
	//_case.CondQueueCase()
	_case.CondQueueCase()
	ctx := context.Background()
	signal.NotifyContext(ctx, os.Kill, os.Interrupt)
}
