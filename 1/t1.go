package main

import (
	"fmt"
	"math"
)

var (
	aa = 1
	bb = "tianya"
)

// 定义变量，只使用默认初始值
func variableZeroValue() {
	var a int
	var s string
	var b bool
	fmt.Println(a, s, b)
}

// 定义变量。赋初值
func variableInitValue() {
	var a, b int = 1, 2
	b = 3
	var s string = "haha"
	fmt.Println(a, b, s)
}

// 类型推断
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

// 最简定义变量方式
func variableShorter() {
	a, b, c, d := 3, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	//variableZeroValue()
	//variableInitValue()
	//variableTypeDeduction()
	//variableShorter()
	triangle()
}