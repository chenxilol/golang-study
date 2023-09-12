package _case

import "fmt"

// 泛型类型

type user struct {
	ID   int
	Name string
	Age  uint8
}
type address struct {
	ID       int
	Province string
	City     string
}
type List[T any] []T
type MapT[k comparable, v any] map[k]v

// 集合转列表
func mapToList[k comparable, v any](mapT MapT[k, v]) List[any] {
	list := make(List[any], 0)
	for k, v := range mapT {
		list = append(list, k)
		list = append(list, v)
	}
	return list
}
func TypeCase() {
	mapCase := make(map[int64]user, 0)
	mapCase[1] = user{
		ID:   2,
		Name: "sdv",
		Age:  18,
	}
	list := mapToList(mapCase)
	fmt.Println(list)
}
