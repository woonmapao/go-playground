package channels

import (
	"fmt"
	"sync"
)

// 1.) A send to a nil channel blocks forever
// 2.) A receive from a nil channel blocks forever

// var c chan s string // c is nil
// c <- "hello"		// block
// fmt.Println(<-c)	// block
//-------------------------------------------------

// A send to a closed channel panics
// A receive from a closed channel returns the zero value immediately

// Function to demonstrate basic channel usage
func BasicChannelExample() {
	fmt.Println("Basic Channel Example:")

	// Create a channel of integers
	intChannel := make(chan int)

	// Start a goroutine to send values to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			intChannel <- i
		}
		close(intChannel) // Close the channel after sending all values
	}()

	// Receive values from the channel
	for value := range intChannel {
		fmt.Println("Received:", value)
	}
}

// Function to demonstrate buffered channels
func BufferedChannelExample() {
	fmt.Println("Buffered Channel Example:")

	// Create a buffered channel of strings with a capacity of 3
	stringChannel := make(chan string, 3)

	// Send values to the buffered channel without blocking
	stringChannel <- "One"
	stringChannel <- "Two"
	stringChannel <- "Three"

	// Close the channel after sending all values
	close(stringChannel)

	// Receive values from the buffered channel
	for value := range stringChannel {
		fmt.Println("Received:", value)
	}
}

// Function to demonstrate synchronization with channels using a WaitGroup
func SynchronizationWithChannels() {
	fmt.Println("Synchronization with Channels Example:")

	// Create a channel to signal completion
	done := make(chan struct{})

	// Use a WaitGroup for synchronization
	var wg sync.WaitGroup

	// Start two goroutines
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Start\n", id)
			// Do some work
			fmt.Printf("Goroutine %d: End\n", id)
		}(i)
	}

	// Start a goroutine to close the done channel when all other goroutines are done
	go func() {
		wg.Wait()   // Wait for all goroutines to finish
		close(done) // Close the done channel to signal completion
	}()

	// Block until the done channel is closed
	<-done
	fmt.Println("All Goroutines Completed")
}
