package mutexes

import (
	"fmt"
	"sync"
	"time"
)

// Shared counter variable
var counter int

// Mutex to synchronize access to the shared variable
var mutex sync.Mutex

// Function to demonstrate basic usage of a mutex
func BasicMutexExample() {
	fmt.Println("Basic Mutex Example:")

	// Use a WaitGroup for synchronization
	var wg sync.WaitGroup

	// Start two goroutines incrementing the counter
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 3; j++ {
				incrementCounter(id, j)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter Value:", counter)
}

// Function to increment the shared counter with a mutex
func incrementCounter(goroutineID, iteration int) {
	// Lock the mutex to synchronize access to the shared variable
	mutex.Lock()
	defer mutex.Unlock()

	// Simulate some work
	time.Sleep(time.Millisecond)

	// Increment the counter
	counter++

	fmt.Printf("Goroutine %d, Iteration %d: Counter = %d\n", goroutineID, iteration, counter)
}
