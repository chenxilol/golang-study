package main

import (
	"context"
	"demo/grpc/hello_grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"sync"
	"time"
)

func main() {
	addr := ":8080"
	// 使用grpc_Dial 创建一个到指定地址的grpc连接
	// 此处使用不安全的证书连接，未实u现 SSl /TLS 连接
	//  grpc.WithTransportCredentials(credentials.NewClientTLSFromCert()) 用于创建一个安全的连接
	dial, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dial.Close()
	// NewHelloServiceClient 用于创建客户端对象

	client := hello_grpc.NewEchoClient(dial)
	UnaryAPI(client)
	//serverStream(client)
	//clientStream(client)
	bidirectionalStream(client)
}

// UnaryAPI 单一实例的响应
func UnaryAPI(client hello_grpc.EchoClient) {
	hello, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "chenxi",
		Message: "",
	})
	num, err := client.AddNum(context.Background(), &hello_grpc.AddNumRequest{
		A: 4,
		B: 10,
	})
	fmt.Println(num)
	if err != nil {
		return
	}
	fmt.Println(hello)
}
func serverStream(client hello_grpc.EchoClient) {
	stream, err := client.ServerStreamingEcho(context.Background(), &hello_grpc.HelloRequest{
		Name:    "chenxi",
		Message: "hello grpc",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		recv, err := stream.Recv()
		if err != nil && err != io.EOF {
			log.Fatal(err)
			return
		}
		if err == io.EOF {
			log.Printf("全部读取完毕")
			break
		}
		log.Printf("Recv data %s %s", recv.GetName(), recv.GetMessage())
	}
}
func clientStream(client hello_grpc.EchoClient) {
	stream, err := client.ClientStreamingEcho(context.Background())
	if err != nil {
		log.Printf("Sum() error :%v", err)
		return
	}
	for i := 0; i < 10; i++ {
		err := stream.Send(&hello_grpc.HelloRequest{
			Name:    fmt.Sprintf("这是客户端发送的第 %d 请求", i),
			Message: "success 上岸",
		})
		if err != nil {
			log.Printf("send error %v", err)
			continue
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("server %s %s", resp.GetName(), resp.GetMessage())
}

// bidirectionalStream 双向流
/*
1. 建立连接 获取client
2. 通过client获取stream
3. 开两个goroutine 分别用于Recv()和Send()
	3.1 一直Recv()到err==io.EOF(即服务端关闭stream)
	3.2 Send()则由自己控制
4. 发送完毕调用 stream.CloseSend()关闭stream 必须调用关闭 否则Server会一直尝试接收数据 一直报错...
*/
func bidirectionalStream(client hello_grpc.EchoClient) {
	var wg sync.WaitGroup
	stream, err := client.BidirectionalStreamingEcho(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Server Closed")
				break
			}
			if err != nil {
				continue
			}
			fmt.Printf("Recv Data :%v", req.GetMessage())
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			err := stream.Send(&hello_grpc.HelloRequest{
				Name:    fmt.Sprintf("这是客户端发送的第 %d 次", i),
				Message: "hello da shuai bi",
			})
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Printf("Send error%v", err)
			return
		}
	}()
	wg.Wait()
}
