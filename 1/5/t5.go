package main

import (
	"fmt"
)

func div(a, b int) (int, int) {
	return a / b, a % b
}

// 函数(传入函数名--传入值定义--传入函数类型，函数传入值) 函数类型
func apply(op func(int, int) int, a,b int) int {
	return op(a, b)
}

// 可变长参数
func sum2(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

// 指针
func swap_by_value(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	// q, _ := div(1, 3)
	// fmt.Println(q)
	// result := apply(func(x int, y int) int {
	// 	return x + y
	// }, 10, 4)
	// fmt.Println(result)
	// sum2(1, 2, 3, 4)

	a, b := 3, 4
	swap_by_value(&a, &b)
	fmt.Println(a, b)
}

