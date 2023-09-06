package main

import (
	"context"
	_case "context-text/case"
	"os"
	"os/signal"
)

func main() {
	_case.ContextCase()

	//通过通道关闭
	//sig := make(chan os.Signal)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//<-sig

	// 通过context关闭
	ctx := context.Background()

	signal.NotifyContext(ctx, os.Kill, os.Interrupt)
}
