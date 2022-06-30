package main

import (
	"fmt"
)

func unbufferedAndBuffered() {
	// Unbuffered channels are created without specifying a size.
	unbufferedChannel := make(chan string)

	// This function is an immediately executed go routine.
	go func() {
		// chan <-
		// when the arrow points to the end, this is a send.
		unbufferedChannel <- "ping"
	}()

	// <- chan
	// When the arrow comes first the functionality is reversed
	// You better believe that this a receive.
	msg := <-unbufferedChannel
	fmt.Println(msg)

	// Unbuffered channels will block if they are not read from.
	// if instead of using a go routine, a regular function was used
	//   func() { chan <- "ping" }()
	// the channel will never be ready to be written to as the func
	// is executed on the main execution thread. The writing is blocked
	// so the read will never be hit.

	// a buffered channel is created by specifying a size
	bufferedChannel := make(chan string, 1)
	// because the channel is buffered it can be written to as it
	// can accept 1 value before it blocks.
	func() { bufferedChannel <- "pong" }()
	msg = <-bufferedChannel

	fmt.Println(msg)

}
