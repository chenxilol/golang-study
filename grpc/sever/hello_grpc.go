package main

import (
	"context"
	"demo/grpc/hello_grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloGrpc struct {
}

func (HelloGrpc) SayHello(ctx context.Context, req *hello_grpc.HelloRequest) (res *hello_grpc.HelloResponse, err error) {
	fmt.Println(req)
	return &hello_grpc.HelloResponse{
		Name:    "ChenxiLol",
		Message: "进军grpc",
	}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// 先初始化grpc
	s := grpc.NewServer()
	sever := HelloGrpc{}
	hello_grpc.RegisterHelloServiceServer(s, &sever)
	fmt.Printf("grpc sever start ：8080")
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}
