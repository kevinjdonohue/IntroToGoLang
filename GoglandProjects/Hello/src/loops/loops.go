package loops

import (
	"fmt"
	"strconv"
)

func PrintMessageFor(message string, numberOfTimes int) {
	for i := 0; i < numberOfTimes; i++ {
		fmt.Println("(" + strconv.Itoa(i + 1) + ") " + message)
	}
}

func PrintMessageWhile(message string, numberOfTimes int) {
	i := 0
	for i < numberOfTimes {
		fmt.Println("(" + strconv.Itoa(i + 1) + ") " + message)
		i++
	}
}

func PrintMessageBreak(message string, numberOfTimes int) {
	i := 0

	for {
		if i >= numberOfTimes {
			break
		}

		fmt.Println("(" + strconv.Itoa(i + 1) + ") " + message)
		i++
	}
}

func PrintMessageContinue(message string, numberOfTimes int) {
	i := 0

	for {
		if i >= numberOfTimes {
			break
		}

		if i % 2 == 0 {
			i++
			continue
		}

		fmt.Println("(" + strconv.Itoa(i + 1) + ") " + message)
		i++
	}
}

func PrintMessageSlice(messages []string) {
	//NOTE:  You can choose to not use the index value by doing: _, msg
	//If you only need the index value you can just do i; you don't have to do i, _
	// If you do i, msg then you must use both values; not using a variable is an error

	for i, msg := range messages {
		fmt.Println("(" + strconv.Itoa(i + 1) + ") " + msg)
	}
}

func PrintCharactersFromMessage(message string) {

	for pos, char := range message {
		fmt.Printf("Character: %#U starts at byte position %d\n", char, pos)
	}
}
