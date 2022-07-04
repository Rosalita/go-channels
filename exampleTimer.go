package main

import (
	"fmt"
	"time"
)

func timer() {
	// time.NewTimer() provides a channel which will be notified
	// after the specified time has passed.
	timer1 := time.NewTimer(2* time.Second)
	<- timer1.C // This is a blocking receive to wait for channel notification.
	fmt.Println("timer1 fired")

	// It is possible to stop a timer before it fires.
	timer2 := time.NewTimer(time.Second)

	// This goroutine will print when the timer fires
	go func(){
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	// instead of allowing the timer to fire, it is stopped
	isStopped := timer2.Stop()
	if isStopped {
		fmt.Println("timer2 was stopped")
	}

	// This sleep is just to illustrate the timer didnt
	// fire when its notification time was reached.
	time.Sleep(2 *time.Second)
}