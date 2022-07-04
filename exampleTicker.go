package main

import (
	"fmt"
	"time"
)

func ticker() {
	// Tickers are for when you want to do something
	// repeatedly at regular intervals.

	ticker := time.NewTicker(500 *time.Millisecond) 
	done := make(chan bool)

	// Start a new go routine
	go func(){
		for{
			// This is a blocking select statement as it has no default case.
			// It will block until one of the cases is met.
			select { 
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at: ", t)
			}
		}
	}()

	// Sleep long enough to let the ticker tick three times.
	time.Sleep(1600 * time.Millisecond)
	// tickers can be stopped in the same way as timers.
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}