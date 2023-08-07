package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerId   string
	ServerPort int
	ServerName string
	conn       net.Conn
}

func NewClient(id string, port int) *Client {
	client := Client{
		ServerId:   id,
		ServerPort: port,
	}
	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", id, port))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	client.conn = conn
	return &client
}
func a() {
	client := NewClient("127.0.0,1", 8888)
	if client != nil {
		fmt.Println(client)
	}
	fmt.Println("链接服务成功")
	select {}
}
