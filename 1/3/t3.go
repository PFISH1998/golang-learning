package main

import (
	"io/ioutil"
	"fmt"
)

func readFile() {
	const filename = "abc.txt"
	// go函数可以返回两个值
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		// content 是  []bytes, 用%s 可以打印
		fmt.Println(string(contents))
	}
	fmt.Printf("%s", contents)
}

func main() {
	readFile()
}