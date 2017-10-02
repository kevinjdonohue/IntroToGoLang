package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type printer func(string) ()

func Greet(salutation Salutation, do printer, isFormal bool) {
	message, alternate := createMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	switch name {
		case "Bob": prefix = "Mr "
		case "Joe": prefix = "Dr "
		case "Mary": prefix = "Mrs "
		default: prefix = "Dude "
	}

	return
}

func CreatePrintFunction(custom string) printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func createMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = greeting + " " + name
	return
}