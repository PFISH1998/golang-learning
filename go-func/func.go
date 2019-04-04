package main

import "fmt"

type user struct {
	name  string
	email string
}

type test struct {
	name string
	email string
}

type admin struct {
	user
	level string
}

// 普通函数定义 func 方法名 （入参） 返回值 {}
// 自定义类型的函数定义 func （接收者） 方法名 （入参） 返回值 {}
// 值传递，拷贝一份 user
func (u user) notify() {
	fmt.Println("pass-by-value", u.name, u.email)
	u.email = "0@qq.com"
}

// 传递地址，会改变原来的
func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

func main() {
	bill := user{"bill", "123@qq.com"}
	bill.notify()
	fmt.Println("1", bill.email)

	// 指向 user 类型的指针
	lisa := &user{"lisa", "qwe@qq.com"}
	lisa.notify()
	fmt.Println("2", lisa.email)

	bill.changeEmail("bill@qq.com")
	fmt.Println("3", bill.email)

	// 嵌入类型， 嵌入类型只需要声明这个类型的名字就可以了
	ad := admin {
		user: user{"nn", "111@qq.com"},
		level: "super",
	}
	// 访问内部类型的方法
	ad.user.notify()
	// 直接从外部访问
	ad.notify()

}