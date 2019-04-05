package main
/*
defer 的调用机制是 将defer 语句压入栈中，
当函数结束时（包括正常执行结束 / return / panic 出错结束等） 从栈中依次执行 defer
*/

import (
	"fmt"
	"os"
	"bufio"
	"errors"
)

func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	fmt.Fprint(write, "123abc")
	// 当函数结束时， 从 defer 栈中执行语句-后进先出
}

func recove() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err)
		} else {
			panic(errors.New("not konw error"))
		}
	}()

	b := 0
	a := 5/b
	fmt.Println(a)
}

func main() {
	recove()
	writeFile("./defer.txt")
}