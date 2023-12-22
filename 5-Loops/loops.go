package loops

import "fmt"

// Function to demonstrate for loop
func ForLoopExample() {
	fmt.Println("For Loop Example:")

	// Basic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// For loop with range for iterating over slices or arrays
	numbers := []int{2, 4, 6, 8, 10}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

// Function to demonstrate while-like loop using for in Go
func WhileLoopExample() {
	fmt.Println("While-like Loop Example:")

	// While-like loop
	counter := 0
	for counter < 3 {
		fmt.Println("Count:", counter)
		counter++
	}
}

// Function to demonstrate infinite loop with break statement
func InfiniteLoopExample() {
	fmt.Println("Infinite Loop Example:")

	// Infinite loop with break statement
	i := 1
	for {
		fmt.Println("Iteration:", i)
		i++

		if i > 3 {
			break
		}
	}
}

// Function to demonstrate continue statement in a loop
func ContinueStatementExample() {
	fmt.Println("Continue Statement Example:")

	// Loop with continue statement
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println(i)
	}
}
