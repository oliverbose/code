package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func main() {
	bill := user{
		name:  "bill",
		email: "bill@olddomain.com",
	}
	bill.notify()
	bill.changeEmail("bill@newdomain.com")
	bill.notify()
	(&bill).changeEmail("bill@newerdomain.com")
	bill.notify()

	lisa := &user{"Lisa", "lisa@lisa.com"}
	lisa.notify()
	(*lisa).notify()

}

func (u user) notify() {
	fmt.Printf("Sending message to %s <%s>\n", u.name, u.email)
}

func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}
