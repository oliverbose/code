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
type process struct {
	id   int
	name string
}

func (u user) notify() {
	fmt.Printf("Sending message to user %s <%s>\n", u.name, u.email)
}

func (p process) notify() {
	fmt.Printf("Sending message to process %v <%s>\n", p.id, p.name)
}

func main() {
	bill := user{"bill", "bill@email.com"}
	proc := process{322, "myapp"}
	sendNotification(bill)
	sendNotification(proc)
}

func sendNotification(n notifier) {
	n.notify()
}
