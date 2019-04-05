package main

import "fmt"

func defineArray() [3]int {
	// 定义数组，不赋初值
	var arr1 [5]int
	// 定义数组，赋初值
	arr2 := [3]int{1,2,3}
	// 由编译器计算长度
	arr3 := [...]int{1,2,3,4,5}
	// 创建数组指针
	arr4 := [2]*string{new(string), new(string)}
	*arr4[0] = "hello"
	*arr4[1] = "go"
	// 为索引位置设置值
	arr5 := [3]int{1:10}
	// 二维数组
	var grid [4][5]int
	
	arr6 := arr2

	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6, grid)
	fmt.Println(*arr4[0])

	return arr2
}

// 数组是值传递，入参会拷贝一份数组，使用指针可以实现“引用传递”
func loopArry(arr2 *[3]int)  *[3]int {
	// 通用方法
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	// 最简方法，只获取下标
	for i := range arr2 {
		fmt.Println(arr2[i])
	}
	// 最简方法，获取下标和值
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	// 最简方法，获取值，省略下标
	for _, v := range arr2 {
		fmt.Println(v)
	}
	arr2[2] = 3
	change_arr := arr2
	return change_arr

}

func main() {
	arr := [3]int{1,2,4}
	fmt.Println(arr)
	a := loopArry(&arr)
	fmt.Println(arr, a)
}
