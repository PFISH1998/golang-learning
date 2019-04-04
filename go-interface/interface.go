package main

import "fmt"

type user struct {
	name string
	age  int
}

// notifiter 是一个定义了通知类行为的接口
type notifier interface {
	// 接口方法
	notify()
	notify2()
}

// 实现接口  值接收者
func (u user) notify() {
	u.name = "p"
	fmt.Println("notify", u)
}

// 指针接收者， 可以改变对象
func (u *user) notify2() {
	fmt.Println("notify2", *u)
	u.name = "p"
}

func sendNotification(n notifier) {
	n.notify()
	n.notify2()
}


func main() {
	u := user{"Jack", 18}
	sendNotification(&u)
	fmt.Println("u", u)
}