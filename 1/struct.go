package main

import "fmt"

// user 类
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// admin 类
type admin struct {
	person user
	level  string
}

func main() {
	// 初始化为 0 值
	var bill user
	fmt.Println(bill)
    // 创建 user ，并初始化
	lisa := user {
		name: "lisa",
		email: "123@qq.com",
		ext: 123,
		privileged: true,
	}
	// 直接定义
	lisa2 := user {"lisa2", "345@1163.com", 123, true}
	fmt.Println("lisa", lisa, "lisa2", lisa2)

	// 含有自定义变量的 struct 进行初始化
	fred := admin{
		person: user{"fred", "fred@qq.com", 123, true},
		level: "admin",
	}
	fmt.Println("fred", fred)
}