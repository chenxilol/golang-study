package main

import (
	"context"
	"demo/grpc/hello_grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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
	client := hello_grpc.NewHelloServiceClient(dial)
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
