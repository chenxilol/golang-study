package main

import "fmt"

// 二分查找
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middleIndex := (low + high) / 2
		guess := nums[middleIndex]
		if guess == target {
			return middleIndex
		}
		if guess > target {
			high = middleIndex - 1
		} else {
			low = middleIndex + 1
		}

	}
	return 0
}

func main() {
	arr := []int{1, 2, 5, 8, 11, 13}
	a := binarySearch(arr, 13)
	fmt.Println(a)
}
