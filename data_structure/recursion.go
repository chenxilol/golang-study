package main

import (
	"fmt"
	"math/big"
	"net"
)

// 斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func NewFibonacci() {
	for i := 1; i <= 10; i++ {
		results := fibonacci(i)
		fmt.Println(results)
	}
}

// Item 套娃递归
type Item struct {
	ID       int
	Type     string
	Children *Item
}

// FindDiamond 寻找套娃中的钻石
func FindDiamond(item *Item) *Item {
	if item == nil {
		return nil
	}
	if item.Type == "钻石" {
		return item
	}
	return FindDiamond(item.Children)
}
func NewFindDiamond() {
	item := Item{
		ID:   1,
		Type: "TaoWa",
		Children: &Item{
			ID:   2,
			Type: "TaoWa",
			Children: &Item{
				ID:   3,
				Type: "TaoWa",
				Children: &Item{
					ID:   4,
					Type: "TaoWa",
					Children: &Item{
						ID:   5,
						Type: "钻石",
					},
				},
			},
		},
	}
	fmt.Println(FindDiamond(&item))
}

// 递归实现阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	//fmt.Println(factorial(3))
	fmt.Println(ip2int("1.0.79.255"))
}

// 将ip地址转换为整数
func ip2int(ip string) int64 {
	ips := net.ParseIP(ip).To4()
	if ips == nil {
		fmt.Println("ip地址错误")

	}
	ret := big.NewInt(0)
	num := ret.SetBytes(ips).Int64()
	return num
}
