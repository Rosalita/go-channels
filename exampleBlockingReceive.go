package main

import (
	"fmt"
	"time"
)

func blockingReceive() {
	// An example of using a blocking receive to wait for a goroutine  to finish
	done := make(chan bool, 1) // done is a buffered channel
	go worker(done)            // start the worker and give it the channel

	// read from the done channel
	<-done // if this line was removed, execution would not block and would end before the worker ran
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
