package main

import (
	_case "GolandProjects/sync/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	_case.MutexCase()
	ctx := context.Background()

	signal.NotifyContext(ctx, os.Kill, os.Interrupt)
}
