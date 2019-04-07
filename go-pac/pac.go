package main

import "fmt"

// 闭包
func adder() func(value int) int {
	sum :=0
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {
	add := adder()
	for i := 0; i < 10; i ++ {
		fmt.Printf("0 + ... + %d = %d", i, add(i))
		fmt.Println()
	}
}