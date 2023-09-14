package main

import (
	"context"
	"demo/grpc/hello_grpc"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

type HelloGrpc struct {
	hello_grpc.EchoServer
}

func (HelloGrpc) SayHello(ctx context.Context, req *hello_grpc.HelloRequest) (res *hello_grpc.HelloResponse, err error) {
	fmt.Println(req)
	return &hello_grpc.HelloResponse{
		Name:    "ChenxiLol",
		Message: "进军grpc",
	}, nil
}
func (HelloGrpc) AddNum(ctx context.Context, req *hello_grpc.AddNumRequest) (res *hello_grpc.AddNumResponse, err error) {
	sum := req.A + req.B
	return &hello_grpc.AddNumResponse{
		Sum: sum,
	}, nil
}

// ServerStreamingEcho 服务端流：服务端可以发送多个数据给客户端,注意当使用流传输数据的时候，不可以用指针 stream
func (HelloGrpc) ServerStreamingEcho(req *hello_grpc.HelloRequest, stream hello_grpc.Echo_ServerStreamingEchoServer) error {
	log.Println(req.Name, req.Message)
	for i := 0; i < 10; i++ {
		err := stream.Send(&hello_grpc.HelloResponse{
			Name:    fmt.Sprintf("这是server的第 %d 次响应", i),
			Message: "hello grpc",
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// ClientStreamingEcho 客户端流:客户端可以发送多个数据给服务端
func (HelloGrpc) ClientStreamingEcho(stream hello_grpc.Echo_ClientStreamingEchoServer) error {
	// for 循环接受客户端发送的消息
	for {
		// 通过Recv() 不断获取客户端 seng() 推送的消息
		recv, err := stream.Recv()
		// 这里相当于处理客户端发来的消息
		log.Printf("Recved %v %v", recv.GetName(), recv.GetMessage())
		// err == io.EOF 表示已经获取全部数据
		if err == io.EOF {
			log.Println("成功接受客户端的所有请求数据")
			// sendAndClose 返回并关闭连接
			return stream.SendAndClose(&hello_grpc.HelloResponse{
				Name:    "zhang",
				Message: "666",
			})
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
}

// BidirectionalStreamingEcho 双向流服务器，他的生命周期为
/*
1. 建立连接 获取client
// 2. 通过client调用方法获取stream
// 3. 开两个goroutine（使用 chan 传递数据） 分别用于Recv()和Send()
// 3.1 一直Recv()到err==io.EOF(即客户端关闭stream)
// 3.2 Send()则自己控制什么时候Close 服务端stream没有close 只要跳出循环就算close了。 具体见https://github.com/grpc/grpc-go/issues/444
*/
func (HelloGrpc) BidirectionalStreamingEcho(stream hello_grpc.Echo_BidirectionalStreamingEchoServer) error {
	var (
		waitGroup sync.WaitGroup
		msgCh     = make(chan string)
	)
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		for msg := range msgCh {
			err := stream.Send(&hello_grpc.HelloResponse{
				Name:    "BidirectionalStreamingEcho",
				Message: msg,
			})
			if err != nil {
				log.Fatal(err)
				return
			}

		}
	}()
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				log.Println("完成数据的读取")
				break
			}
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("Recved :%v \n", req.GetMessage())
			msgCh <- req.GetMessage()
		}
		close(msgCh)
	}()
	waitGroup.Wait()
	return nil
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// 先初始化grpc
	grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		method = "/protos." + strings.TrimPrefix(method, "/pb.")
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			return errors.New(fmt.Sprintf("GRPC 调用失败,method:%s,err:%v", method, err))
		}
		return nil
	})
	s := grpc.NewServer()
	sever := HelloGrpc{}

	hello_grpc.RegisterEchoServer(s, &sever)
	fmt.Printf("grpc sever start ：8080")
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}
