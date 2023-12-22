package mutexes

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Shared data structure
type Data struct {
	value int
}

// Shared data protected by a read-write mutex
var sharedData Data
var rwMutex sync.RWMutex

// Function to demonstrate read and write operations with read-write mutex
func ReadWriteMutexExample() {
	fmt.Println("Read-Write Mutex Example:")

	// Use a WaitGroup for synchronization
	var wg sync.WaitGroup

	// Start multiple goroutines for reading
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			readData(readerID)
		}(i)
	}

	// Start a goroutine for writing
	wg.Add(1)
	go func() {
		defer wg.Done()
		writeData()
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}

// Function for reading data with read-write mutex
func readData(readerID int) {
	// Simulate some work
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Lock for reading (multiple readers allowed)
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	fmt.Printf("Reader %d: Reading Data - Value: %d\n", readerID, sharedData.value)
}

// Function for writing data with read-write mutex
func writeData() {
	// Simulate some work
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Lock for writing (only one writer allowed)
	rwMutex.Lock()
	defer rwMutex.Unlock()

	// Update the shared data
	newValue := rand.Intn(100)
	sharedData.value = newValue

	fmt.Printf("Writer: Writing Data - New Value: %d\n", newValue)
}
