package main

import (
	"fmt"
)

type user struct {
	name   string
	email  string
	ext    int
	active bool
}

type admin struct {
	person user
	level  string
}

type duratuion int64

func main() {
	//emptyUser()
	//literalUser()
	//literalUserShort()
	//makeAdmin()
	makeDuration()
}

func emptyUser() {
	var bill user
	printUser(bill)
}

func literalUser() {
	bill := user{
		name:   "bill",
		email:  "bill@bill.com",
		ext:    123,
		active: true,
	}
	printUser(bill)
}

func literalUserShort() {
	lisa := user{"Lisa", "lisa@lisa.com", 123, true}
	printUser(lisa)
}

func makeAdmin() {
	john := admin{
		person: user{
			name:   "john",
			email:  "john@john.com",
			ext:    432,
			active: true,
		},
		level: "domain",
	}
	printAdmin(john)
}

func makeDuration() {
	var dur duratuion
	dur = duratuion(1000)
	fmt.Printf("Duration: %d", dur)
}
func printUser(aUser user) {
	fmt.Printf("Name: %s, email: %s, ext: %v, active: %v", aUser.name, aUser.email, aUser.ext, aUser.active)
}

func printAdmin(anAdmin admin) {
	fmt.Printf("Name: %s, email: %s, ext: %v, active: %v, level: %s", anAdmin.person.name, anAdmin.person.email, anAdmin.person.ext, anAdmin.person.active, anAdmin.level)
}
