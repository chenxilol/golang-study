package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func initRedis() (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
func main() {
	cli, err := initRedis()
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	cli.Set(ctx, "demo1", "1", 0)
	cli.LPush(ctx, "ListDemo", "1", "2", "3")
	i := cli.Get(ctx, "demo1")
	fmt.Println(i)
	cli.Append(ctx, "demo1", "2")
	j := cli.Get(ctx, "demo1")
	fmt.Println(j)
	cli.Incr(ctx, "demo1")
	cli.Del(ctx, "dem1")
	rdb.Publish(ctx, "channel1", "hello1")
}
