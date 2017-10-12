package main

import (
	"./greeting"
	"fmt"
)

func main() {

	salutations := greeting.Salutations {
		{"Joe", "Hi"},
		{"Mary", "Hello"},
	}

	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	salutations.Greet(greeting.CreatePrintFunction("!!!"), true)
}