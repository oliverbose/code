package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}
type admin struct {
	user
	name  string
	level string
}

func (u *user) notify() {
	fmt.Printf("Sending message to user %s <%s>\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("Sending message to admin %v <%s>\n", a.name, a.level)
}

func main() {
	adm := admin{user{"adm", "adm@email.com"}, "adm2", "domain"}
	sendNotification(&adm.user)
	sendNotification(&adm)
}

func sendNotification(n notifier) {
	n.notify()
}
