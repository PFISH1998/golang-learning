package main

import (
	"fmt"
	"math"
)

func consts() {
	// 指定类型
	const filename string = "filename-const"
	// 不指定类型，表示类型不定
	const a, b = 3, 4
	var c int

	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, a, b, c)

}

func enums() {
	// 使用const块来实现枚举
	const (
		java = 0
		cpp  = 1
		c    = 2 
	)
	fmt.Println(java, cpp, c)

	// 使用iota实现自增枚举
	const (
		java1 = iota
		cpp1
		c1
	)
	fmt.Println(java1, cpp1, c1)
}

func main() {
	consts()
	enums()
}