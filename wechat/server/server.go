package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP   string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:   ip,
		Port: port,
	}
}

// 启动服务器的接口
func (s *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("服务器启动成功，等待客户端连接...")
	defer listen.Close()
	for {
		// 等待客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			continue
		}
		// 启动一个协程和客户端保持通讯
		go func() {
			for {
				// 读取客户端发送的数据
				buf := make([]byte, 4096)
				
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("conn.Read err:", err)
					return
				}
				// 显示客户端发送的内容到服务器的终端
				fmt.Println("服务器读到的数据:", string(buf[:n]))
			}
		}()
	}
}
func main() {
	NewServer("127.0.0.1", 8888).Start()
}
