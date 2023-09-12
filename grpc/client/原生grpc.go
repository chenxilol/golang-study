package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}
type Res struct {
	Num int
}

func main() {
	req := Req{1, 9}
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	var res Res
	// 这是一个远程方法调用的操作，通常用于客户端通过网络调用远程服务器上的方法
	client.Call("Server.Add", req, &res)
	fmt.Println(res)
}
