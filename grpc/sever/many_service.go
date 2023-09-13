package main

import (
	"context"
	"demo/grpc/many_service_grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ByteSort struct {
}

func (ByteSort) Sort(ctx context.Context, request *many_service_grpc.BytesRequest) (response *many_service_grpc.BytesResponse, err error) {
	arr := request.Data
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 提前退出标志，如果一轮遍历没有发生交换，说明已经排序完成
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换arr[j]和arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// 如果一轮遍历没有发生交换，说明已经排序完成
		if !swapped {
			break
		}
	}
	return &many_service_grpc.BytesResponse{DataSort: arr}, err
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	m := ByteSort{}
	many_service_grpc.RegisterManyServiceOneServer(s, &m)
	err = s.Serve(listen)
	if err != nil {
		return
	}
	fmt.Println("grpc start ...")
}
