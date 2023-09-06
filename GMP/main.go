package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	create, err := os.Create("trace.out")
	if err != nil {
		return
	}
	defer create.Close()
	err = trace.Start(create)
	if err != nil {
		return
	}
	fmt.Println("hello GMP")
	trace.Stop()
}
