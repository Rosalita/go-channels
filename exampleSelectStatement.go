package main

import (
	"fmt"
	"time"
)

func selectStatement() {
	// A select statement combines goroutines and channels.
	// It allows code to wait for multiple channel operations

	// Create two new channels.
	ch1 := make(chan string)
	ch2 := make(chan string)

	// In a new thread, wait 1 second then send to channel 1.
	go func() {
		time.Sleep(time.Second)
		ch1 <- "one"
	}()

	// In another new thread, wait 2 seconds then send to channel 2.
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// Select can be used to await both of these values simulataneously
	// and print each one as it arrives.
	// The for loop will select twice
	for i := 0; i < 2; i++ {
		select { // like a switch statement, when the case is true, it executes
		case msg1 := <-ch1:
			fmt.Println("received from channel 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("received from channel 2:", msg2)
		}
		// As this select statement has no default, it wil block until one
		// of its cases is triggered.
	}

	// Two messages are sent to the channels by go routines
	// so if the select statement was called three times this
	// would result in a deadlock because all the go routines would be asleep.
}
