package main

import (
	"fmt"
	"os"
	"bufio"
)

func sum(j int) int {
	var value int
	for i := 0; i <= j; i ++ {
		value += i
	}
	return value
}

// 等同于 while(true)
func deadloop() {
	for {
		fmt.Println("this is deadloop")
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	// 如果出错，结束进程
	if err != nil {
		panic(err)
	}
	// 获取读取器
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	i := sum(100)
	fmt.Println(i)
	printFile("abc.txt")
	// deadloop()
}