package main

import (
	"fmt"
)

func channelDirections() {
	// When specifying function params it is possible to state that channel is only for sending or receiving.
	// The function signature uses chan<- for channels that can only send
	// and <-chan for channels that can only receive.
	// the type of the channel must also be specified

	myChannel := make(chan string, 1)
	send(myChannel, "hello")
	receive(myChannel)
}

// send a message to a channel
func send(ch chan<- string, msg string) {
	ch <- msg
}

// receive a message from a channel
func receive(ch <-chan string) {
	received := <-ch
	fmt.Println("received:", received)
}
