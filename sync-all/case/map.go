package _case

import (
	"fmt"
	"sync"
)

func MapCase() {
	// 读多写少
	mp := sync.Map{}
	mp.Store("name", "chenXi")
	mp.Store("email", "@")
	// 通过key值获取value，如果不存在则返回nil，ok，则返回false
	fmt.Println(mp.Load("name"))
	fmt.Println(mp.Load("email"))
	// 通过key值获取value，如果不存在则设置指定的value并返回
	// ok 为true表示key值存在并返回值，为false表示key 不存在并设置后返回
	fmt.Println(mp.LoadOrStore("hobby", "跑步"))
	fmt.Println(mp.LoadOrStore("hobby", "打篮球"))

	// 根据key获取value后删除该key
	// ok为true表示key存在，为false表示key值不存在
	fmt.Println(mp.LoadAndDelete("hobby"))
	fmt.Println(mp.LoadAndDelete("hobby"))

	// 为集合设置迭代函数，将为集合中的每一个键值对顺序调用该函数，如果该函数返回false，则停止迭代
	// 为遍历集合中所有元素提供遍历
	mp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
