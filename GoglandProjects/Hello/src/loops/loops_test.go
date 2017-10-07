package loops_test

import (
	"testing"
	"loops"
)

func TestPrintMessageFor(t *testing.T) {

	//act
	loops.PrintMessageFor("Hello For", 5)
}

func TestPrintMessageWhile(t *testing.T) {

	//act
	loops.PrintMessageWhile("Hello While", 5)
}

func TestPrintMessageBreak(t *testing.T) {

	//act
	loops.PrintMessageBreak("Hello Break", 5)
}

func TestPrintMessageContinue(t *testing.T) {

	//act
	loops.PrintMessageContinue("Hello Continue", 5)
}

func TestPrintMessageRange(t *testing.T) {

	//arrange
	var messages []string = []string{"Hello Kevin", "Hello Kim", "Hello Kate"}

	//act
	loops.PrintMessageRange(messages)
}

func TestPrintCharactersFromAString(t *testing.T) {

	//arrange
	var message string = "Hello World"

	//act
	loops.PrintCharactersFromAString(message)
}
