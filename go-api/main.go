package main

import (
	"./api"
	"./impl"
)

func sendNotification(n api.Notifiter) {
	n.Notify()
}

func main() {
	u := impl.User{"nana"}
	sendNotification(&u)

	a := impl.Admin{u, "zhao"}
	sendNotification(&a)
}