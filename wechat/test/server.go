package main

import (
	"fmt"
	"net"
	"sync"
)

type Sever struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	SyncMap   sync.RWMutex
	Msg       chan string
}

func NewSever(ip string, port int) *Sever {
	return &Sever{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Msg:       make(chan string),
	}
}
func (this *Sever) Star() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(listen)
	go this.Lis()
	for {
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		go this.Handle(conn)
	}

}
func (this *Sever) Handle(c net.Conn) {
	fmt.Println("链接成功")
	user := Lister(c)
	this.SyncMap.Lock()
	this.OnlineMap[user.Name] = user
	this.SyncMap.Unlock()
	this.BroadCast(user, "已上线")
}

// 像所有用户发送上线信息

func (this *Sever) Lis() {
	for {

		msg := <-this.Msg
		// 将msg发送给全部的在线user
		this.SyncMap.Lock()
		for _, cli := range this.OnlineMap {
			cli.msg <- msg
		}
		this.SyncMap.Unlock()
	}
}
func (this *Sever) BroadCast(user *User, msg string) {
	Msg := user.Addr + msg
	this.Msg <- Msg
}
