package main

import (
	"fmt"
	"time"
)

func poolWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job * 2
	}
}

func workerPool() {

	const numberOfJobs = 5
	// create buffered channels
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)

	//start up three workers
	for i := 0; i < 3; i++ {
		go poolWorker(i, jobs, results)
	}

	// Send some jobs to the jobs channel
	for i := 0; i < numberOfJobs; i++ {
		jobs <- i
	}

	// close the jobs channel to indicate no more jobs
	close(jobs)

	// receive each result from the jobs channel
	for i := 0; i < numberOfJobs; i++ {
		//Note that there are no order guarantees for receiving results
		fmt.Println("result is ", <-results)
	}

}
