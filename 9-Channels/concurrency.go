package main

import (
	"fmt"
	"time"
)

func main() {

	// 1.) A send to a nil channel blocks forever
	// 2.) A receive from a nil channel blocks forever

	// var c chan s string // c is nil
	// c <- "hello"		// block
	// fmt.Println(<-c)	// block
	//-------------------------------------------------

	// A send to a closed channel panics
	// A receive from a closed channel returns the zero value immediately

}

type email struct {
	message string
	date    time.Time
}

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go func() {

		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()

	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)

}

func countReports(numSentCh chan int) int {
	total := 0
	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent

	}
	return total
}

func concurrentFib(n int) {
	chInts := make(chan int)
	x, y := 0, 1
	go func() {
		for i := 0; i < n; i++ {
			chInts <- x
			x, y = y, x+y
		}
		close(chInts)
	}()

	for item := range chInts {
		fmt.Println(item)
	}
}

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case _, ok := <-chEmails:
			if !ok {
				return
			}
			// logEmail

		case _, ok := <-chSms:
			if !ok {
				return
			}
			// logSms
		}
	}
}

// The 'default' case in a 'select' statement executes
// immediately if no other channel has a value ready

// * A 'default' case stop the 'select' statement from blocking

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {

	for {
		select {

		case <-snapshotTicker:
			//takeSnapshot()

		case <-saveAfter:
			//saveSnapshot()
			return

		default:
			//waitForData()
			time.Sleep(time.Millisecond * 500)
			fmt.Println("waiting for data")
		}

	}
}
