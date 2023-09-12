package _case

import "fmt"

func Sum[T CusNumT](a, b T) T {
	return a + b
}

type CusNumT interface {
	uint8 | int32 | float64 | ~int64
}

func SimpleCase() {

	// 衍生类型
	type MyInt64 int64

	// 别名
	type MyInt32 = int32
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println(Sum(a, b))
	fmt.Println(Sum(a1, b1))
	var a3, b3 MyInt64 = 3, 4
	fmt.Println(Sum(a3, b3))
	d, c := "zhang", "zhang"
	fmt.Println(getBuildInComparable(d, c))
}

// 内置类型，一个comparable 主要是用于比较，一个any类型可以传入任何类型
func getBuildInComparable[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}
