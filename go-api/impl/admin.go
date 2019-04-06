package impl

import "fmt"

type Admin struct {
	// 嵌入类型
	User
	Name string
}

func (a *Admin) Notify()  {
	fmt.Println("admin", *a)
}