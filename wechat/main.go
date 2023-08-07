package main

import (
	"fmt"
	"strings"
)

func main() {
	//客户端
	y := "to|user"
	test := strings.Split(y, "|")[0]
	fmt.Println(test)
	s := NewServer("127.0.0.1", 8888)
	s.Start()

}
