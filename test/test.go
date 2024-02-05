package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	myctxErr := fmt.Errorf("自定义ctxErr错误")
	ctx, cancel := context.WithTimeoutCause(context.Background(), 5*time.Second, myctxErr)
	context.AfterFunc(ctx, func1)
	defer cancel()
	time.Sleep(6 * time.Second)
	fmt.Println(context.Cause(ctx))
}

// evictEvent
func func1() {
	fmt.Println(666)
}
