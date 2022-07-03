package main

import "fmt"

func rangeOverChannel() {

	queue := make(chan string, 2) //queue is a buffered channel.

	// Add some values to the queue.
	queue <- "one"
	queue <- "two"
	close(queue) // It is possible to close a non-empty queue.

	// range iterates over each item as it is received from the channel.
	for item := range queue { 
		// Even if a channel is closed, items will be received.
		fmt.Println("Next queue item is: ", item)
	}
}
