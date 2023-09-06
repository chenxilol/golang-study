package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP        string
	Port      int
	OnlineMap map[string]*User
	SyncMap   sync.RWMutex
	Message   chan string
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

// Handler 消息中间件
func (s *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	user := NewUser(conn, s)
	// 用户上线，将用户加入到online-map中
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		for {
			count, err := conn.Read(buf)
			if err != nil {
				user.Offline()
				return
			}
			// 因为字节最后一个字符为EOF
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			}
			msg := string(buf[:count-1])
			user.DoMessage(msg)
			// 用户的任意消息，代表当前用户是一个活跃的用户
			isLive <- true
		}
	}()
	for {
		select {
		// 监听channel上的数据流动
		case <-isLive:
			// 当前用户是活跃的，应该重置定时器
			// 不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(time.Second * 600):
			// 已经超时
			close(user.C)
			_ = conn.Close()
			return

		}
	}
}

// ListenMessage 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message
		// 将msg发送给全部的在线user
		s.SyncMap.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.SyncMap.Unlock()
	}
}

// BroadCast 广播消息的方法
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Message <- sendMsg
}

// Start 启动服务器的接口
func (s *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("服务器启动成功，等待客户端连接...")
	defer listen.Close()
	// 启动监听Message的goroutine
	go s.ListenMessage()
	for {
		// 等待客户端连接，只有链接成功才会返回conn
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			continue
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println("conn.Close err:", err)
				return
			}
		}(conn)
		// 启动一个协程和客户端保持通讯
		go s.Handler(conn)
	}
	// 广播用户上线信息
}
