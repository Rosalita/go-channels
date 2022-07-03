package main

import "fmt"

func closeChannels() {
	// Closing a channel indicates that no more values will be sent to it.
	// This can be useful to communicate completion.

	jobs := make(chan int, 5) // buffered channel
	done := make(chan bool)   // unbuffered channel

	go func() {
		// Using a for statement to receive
		// The second value, more, will be false if the channel has been closed.
		for {
			job, more := <-jobs
			if !more { // if there are no more jobs
				fmt.Println("all jobs done")
				done <- true
			}
			fmt.Println("received job", job)
		}
	}()

	// Send three jobs.
	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}

	close(jobs) // Then close the jobs channel.
	fmt.Println("sent all jobs")

	<-done // This is a blocking receive which will keep the code running until done is received.
}
