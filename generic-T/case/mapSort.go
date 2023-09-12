package _case

import (
	"fmt"
	"sort"
)

// 创建一个Map集合，可以接受任意值

type MapTSort[k comparable, v any] map[k]v
type List1[t comparable] []t

func MapSort[k comparable, v any](ma MapTSort[k, v]) List1[k] {
	list := make(List1[k], 0)
	for key, _ := range ma {
		list = append(list, key)
	}
	return list
}
func SortMapKey() {
	m := make(map[int]string, 0)
	m[1] = "zhang"
	m[2] = "li"
	m[3] = "wang"
	m[9] = "h"
	m[4] = "a"
	m[100] = "100"
	m[6] = "6"
	list := MapSort(m)
	sort.Ints(list)
	for _, value := range list {
		if v, ok := m[value]; ok {
			fmt.Println(value, v)
		}
	}
}
