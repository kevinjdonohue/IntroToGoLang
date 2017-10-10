package main

import "greeting"

func main() {

	salutations := greeting.Salutations {
		{"Joe", "Hi"},
		{"Mary", "Hello"},
	}

	salutations.Greet(greeting.CreatePrintFunction("!!!"), true)
}