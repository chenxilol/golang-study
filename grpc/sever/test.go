package main

import "fmt"

func bubbleSort(arr []int) {
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
}
func main() {
	s := []int{1, 5, 22, 12}
	bubbleSort(s)
	fmt.Println(s)
}
