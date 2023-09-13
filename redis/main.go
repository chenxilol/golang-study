package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func main() {
	ctx := context.Background()
	err := rdb.Set(ctx, "goRedisTestKey", "goRedisTestValue", 0).Err()
	if err != nil {
		panic(err)
	}
	//err = rdb.HSet(ctx, "goRedisTestHash", "goRedisTestHashKey", "goRedisTestHashValue").Err()
	//val, err := rdb.Get(ctx, "goRedisTestKey").Result()
	//data, err := rdb.HGetAll(ctx, "goRedisTestHash").Result()
	//fmt.Println(val, data)
	//data := make(map[string]interface{})
	//data["goRedisTestHashKey"] = 1
	//data["goRedisTestHashKey2"] = "data"
	//err = rdb.HMSet(ctx, "goRedisTestHash", data).Err()
	//if err != nil {
	//	panic(err)
	//}
	//val := rdb.HGetAll(ctx, "goRedisTestHash").Val()
	//for s, s2 := range val {
	//	fmt.Println(s, s2)
	//}
	//rdb.LPush(ctx, "goRedisTestList", "1", "2", "3", "4", "5")
	//val := rdb.LRange(ctx, "goRedisTestList", 0, -1).Val()
	//fmt.Println(val)

	sub := rdb.Subscribe(ctx, "channel1")
	for ch := range sub.Channel() {
		fmt.Println(ch.Channel)
		fmt.Println(ch.Payload)
	}
}
