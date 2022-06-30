package main

import (
	"fmt"
	"time"
)

func timeout() {
	// Timeouts can be implemented with channels and select.

	// Create a new channel.
	ch1 := make(chan string)

	// Start a new thread that sends to the channel after 2 seconds.

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "not a timeout"
	}()

	// Select statement will execute once for which ever of the cases is true.
	select {
	case msg := <-ch1:
		fmt.Println("received message: ", msg)
	// time.After is a method that accepts a duration and returns
	// a channel which can only receive a time.Time
	// the select statement executes if that channel receives.
	case <-time.After(30 * time.Second):
		fmt.Println("timeout")
	}
}
