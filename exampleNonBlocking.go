package main

import (
	"fmt"
)

func nonBlocking() {
	// Basic sends and receives on channels are blocking.
	// However select can be used with a default case to implement non-blocking selection.

	messages := make(chan string)
	signals := make(chan bool)

	// This is a non-blocking receive.
	select {
	case msg := <-messages: // if there is a value available to receive select will take this case.
		fmt.Println("received message: ", msg)
	default: // if not it will immediately take the default case.
		fmt.Println("no message received")
	}

	msg := "hi"

	// This is a non-blocking send.
	select {
	case messages <- msg: // if msg can be send to the channel select will take this case.
		fmt.Println("sent message: ", msg) // because the messages channel is unbuffered msg cant be sent
	default:
		fmt.Println("no message sent") // so the default case is selected.
	}

	// A select statement can use multiple cases above a default clause to implement 
	// multi-way non-blocking select.
	// This select statement will do non-blocking receives on the messages and signals
	// channels.

	select {
	case msg := <- messages:
		fmt.Println("received message: ", msg)
	case sig := <- signals:
		fmt.Println("received signal: ", sig)
	default:
		fmt.Println("no activity")
	}

}
