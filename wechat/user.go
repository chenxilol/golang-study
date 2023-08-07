package main

import (
	"fmt"
	"net"
	"strings"
)

// User 用户结构体
type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	Server *Server
}

// NewUser 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		Server: server,
	}
	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()
	return user
}

// Online 用户上线业务
func (this *User) Online() {
	this.Server.SyncMap.Lock()
	this.Server.OnlineMap[this.Name] = this
	this.Server.SyncMap.Unlock()
	// 广播当前用户上线消息
	this.DoMessage("已上线")
}

// Offline 用户下线业务
func (this *User) Offline() {
	this.Server.SyncMap.Lock()
	delete(this.Server.OnlineMap, this.Name)
	this.Server.SyncMap.Unlock()
	// 广播当前用户上线消息
	this.DoMessage("已下线")
	defer this.conn.Close()
}

// SendMsg 发送消息
func (this *User) SendMsg(msg string) {
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		return
	}
}

// 处理用户指令
func (this *User) UserMsg(msg string) (s []string) {
	s = strings.Split(msg, "|")
	return
}

// DoMessage 处理消息的业务
func (this *User) DoMessage(msg string) {
	switch {
	case msg == "who":
		this.Server.SyncMap.Lock()
		for _, user := range this.Server.OnlineMap {
			onlineUser := fmt.Sprintf("[%s]%s\n", user.Addr, user.Name)
			this.SendMsg(onlineUser)
		}
		this.Server.SyncMap.Unlock()
	case msg == "offOnline":
		this.Offline()
	case msg == "rename":
		this.SendMsg("请输入新的用户名：")
		buf := make([]byte, 64)
		_, err := this.conn.Read(buf)
		nonZeroData := make([]byte, 0)

		// 遍历原始字节切片，并将非零值添加到新的切片中
		for _, val := range buf {
			if val != 0 {
				nonZeroData = append(nonZeroData, val)
			}
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		this.Server.SyncMap.Lock()
		if _, ok := this.Server.OnlineMap[string(buf[:len(buf)-1])]; ok {
			this.SendMsg("用户名已存在")
			return
		}

		delete(this.Server.OnlineMap, this.Name)
		s := string(nonZeroData[:len(nonZeroData)-1])
		this.Name = s
		fmt.Println(len(s))
		this.Server.OnlineMap[this.Name] = this
		this.Server.SyncMap.Unlock()
		this.SendMsg("您已经更新用户名：" + this.Name)
	case this.UserMsg(msg)[0] == "to":
		s := this.UserMsg(msg)
		this.Server.SyncMap.Lock()
		fmt.Println(len(s[2]))
		if s[1] == "" {
			this.SendMsg("消息格式不正确，请使用to｜用户名|想发送的消息")
		}
		remoteUser, ok := this.Server.OnlineMap[s[1]]
		if !ok {
			this.SendMsg("当前用户不存在")
		}
		remoteUser.SendMsg(fmt.Sprintf("用户%s向你发送  %s  消息\n", this.Name, s[2]))
		this.Server.SyncMap.Unlock()
	default:
		this.Server.BroadCast(this, msg)
		if msg == "offOnline" {
			this.Offline()
		}
	}

}

// ListenMessage 监听当前user channel的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		_, err := this.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
