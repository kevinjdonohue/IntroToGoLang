package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Renamable interface {
	Rename(newName string)
}

func RenameToFred(r Renamable) {
	r.Rename("Fred")
}

func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

type Salutations []Salutation

type Printer func(string) ()

func (salutations Salutations) Greet(do Printer, isFormal bool) {
	for _, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	//shorthand initialization for a map
	prefixMap := map[string]string {
		"Bob": "Mr ",
		"Joe": "Dr ",
		"Amy": "Dr ",
		"Mary": "Mrs ",
	}

	if prefix, exists := prefixMap[name]; exists {
		return prefix
	}

	return "Dude "
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