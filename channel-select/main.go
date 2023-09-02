package main

import (
	_case "channel-select/case"
	"os"
	"os/signal"
)

func main() {
	//_case.Communication()
	//_case.ConCurrentSync()
	_case.NoticeAndMultiplexing()
	ch := make(chan os.Signal, 0)
	// 一旦有channel信号过来证明系统关闭
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
