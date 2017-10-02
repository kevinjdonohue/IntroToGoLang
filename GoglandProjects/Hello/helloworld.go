package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string) ()

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s, CreatePrintFunction("!!!"))
}

func Greet(salutation Salutation, do Printer) {
	message, alternate := CreateMessageVariadic("Kevin", "Hello", "Hi", "Hola", "Hey", "Yo")

	do(message)
	do(alternate)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func CreateMessageVariadic(name string, greeting ...string) (message string, alternate string) {
	message = greeting[0] + " " + name
	alternate = greeting[len(greeting)-1] + " " + name
	return
}
