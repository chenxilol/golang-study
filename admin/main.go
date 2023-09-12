package main

import (
	"demo/admin/core"
	"demo/admin/global"
	"fmt"
)

func main() {
	//core.Viper.Initialization()
	global.Zap = core.InitLoggers()
	defer global.Zap.Sync()
	core.SimpleHttpGet("www.5lmh.com")
	core.SimpleHttpGet("http://www.google.com")
	fmt.Println(global.Config)
}
