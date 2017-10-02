package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Mary", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true)
}