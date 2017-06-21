package main

import (
	"fmt"

	"github.com/goinaction/code/playground/exportIdentifiers/counters"
	"github.com/goinaction/code/playground/exportIdentifiers/entities"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter %v\n", counter)

	user := entities.New("bill", "bill@email.com")
	fmt.Printf("User %s\n", user)
}
