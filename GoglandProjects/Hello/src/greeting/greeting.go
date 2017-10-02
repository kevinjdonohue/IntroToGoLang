package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	switch name {
		case "Bob":
			prefix = "Mr "
		case "Joe", "Amy":
			prefix = "Dr "
		case "Mary":
			prefix = "Mrs "
		default:
			prefix = "Dude "
	}

	return
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func CreateMessage(name string, messagePrefix string) (primaryMessage string, alternateMessage string) {
	primaryMessage = messagePrefix + " " + name
	alternateMessage = "Yo " + " " + name
	return
}