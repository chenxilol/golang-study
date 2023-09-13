package main

import (
	"context"
	"demo/grpc/many_service_grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	client := many_service_grpc.NewManyServiceOneClient(conn)
	datas := []int32{
		232, 34, 1242, 354, 52332, 12, 324, 4,
	}
	sort, err := client.Sort(context.Background(), &many_service_grpc.BytesRequest{Data: datas})
	if err != nil {
		return
	}
	fmt.Println(sort.DataSort)
}
