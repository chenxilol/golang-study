package main

import (
	_case "GolandProjects/sync/case"
	"context"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	//_case.MutexCase()
	_case.MapCase()
	ctx := context.Background()
	runtime.GOMAXPROCS(7)
	signal.NotifyContext(ctx, os.Kill, os.Interrupt)
}
