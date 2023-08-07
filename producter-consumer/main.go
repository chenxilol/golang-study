package main

import (
	"os"
	"os/signal"
	one_one "producter-consumer/core"
	"producter-consumer/out"
	"syscall"
)

func main() {
	// 创建一个data接受chan的管道
	o := out.NewOut()
	// 监听 chan数据
	go o.OutPut()
	//one_one.Exec()
	//one_one.Test()
	//one_one.Exec2()
	one_one.Exec3()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
