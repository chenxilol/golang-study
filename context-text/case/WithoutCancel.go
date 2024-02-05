package _case

import (
	"context"
	"fmt"
	"time"
)

func WithoutCancel() {
	ctx := context.WithoutCancel(context.Background())
	contextCopy, _ := context.WithCancel(ctx)
	go func() {
		select {
		case <-contextCopy.Done():
			fmt.Println("通道关闭")
		default:
			fmt.Println("通道没关闭")
		}
	}()
	time.Sleep(time.Second)
}
