package impl

import "fmt"

// 定义 user
type User struct {
	Name string
}

func (u *User) Notify() {
	fmt.Println("user", *u)
}