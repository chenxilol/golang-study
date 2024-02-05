package main

import (
	"context"
	_case "demo/context-text/case"
	"os"
	"os/signal"
)

func main() {
	_case.ContextCase()
	//_case.WithoutCancel()
	//通过通道关闭
	//sig := make(chan os.Signal)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//<-sig

	// 通过context关闭
	ctx := context.Background()

	ctx, can := signal.NotifyContext(ctx, os.Kill, os.Interrupt)
	<-ctx.Done()
	can()
}
