package main

import (
	"fmt"
	"sync"
	"time"
)

// This is a worker that sleeps to simulate and expensive task.
func waitGroupWorker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func waitGroup() {
	// a waitgroup is used to wait for all the goroutines to finish.
	// if a waitgroup is explicitly passed to a function, pass by reference.
	var wg sync.WaitGroup

	// launch five go routines
	for i := 0; i < 5; i++ {
		// increment the waitgroup counter for each
		wg.Add(1)

		// as each iteration of this loop uses the same variable for i
		// i could have changed by the time the go routine in this loop
		// is executed. To prevent this, we will pin i to a new variable
		id := i

		// and use the new variable, which is not shared with future iterations
		// of this loop to start the go routine.

		go func() {
			defer wg.Done()
			waitGroupWorker(id)
		}()

		// Block until the waitgroup counter goes back to 0.
		// which means all the workers are done.

		wg.Wait()
	}
}
