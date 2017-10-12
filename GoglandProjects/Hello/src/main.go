package main

import (
	"fmt"

	"./greeting"
)

func main() {

	salutations := greeting.Salutations{
		{"Joe", "Hi"},
		{"Mary", "Hello"},
	}

	//fmt.Fprintf(&salutations[0], "The count is %d", 10)

	//done := make(chan bool)
	//
	//go func() {
	//	salutations.Greet(greeting.CreatePrintFunction("<C>"), true)
	//	done <- true
	//}()
	//
	//salutations.Greet(greeting.CreatePrintFunction("!!!"), true)
	//
	//<- done

	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)

	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")

		}
	}

	//for s := range c {
	//	fmt.Println(s.Name)
	//}
}
