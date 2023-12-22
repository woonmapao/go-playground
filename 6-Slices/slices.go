package slices

import "fmt"

// Function to demonstrate basic operations on slices
func BasicSliceOperations() {
	fmt.Println("Basic Slice Operations Example:")

	// Declare and initialize a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", numbers)

	// Accessing elements in a slice
	fmt.Println("First Element:", numbers[0])
	fmt.Println("Last Element:", numbers[len(numbers)-1])

	// Slicing a slice
	subSlice := numbers[1:4]
	fmt.Println("Sub-slice:", subSlice)

	// Modifying a slice
	numbers[2] = 99
	fmt.Println("Modified Slice:", numbers)

	// Appending to a slice
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("After Append:", numbers)
}

// Function to demonstrate iterating over a slice
func IterateOverSlice() {
	fmt.Println("Iterate Over Slice Example:")

	// Declare and initialize a slice
	fruits := []string{"Apple", "Banana", "Orange", "Mango"}

	// Iterate over the slice using range
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
}

// Function to demonstrate the make function for slices
func MakeFunctionForSlices() {
	fmt.Println("Make Function for Slices Example:")

	// Creating a slice with make
	intSlice := make([]int, 3, 5)
	fmt.Println("Initial Slice:", intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	// Appending elements to the slice
	intSlice = append(intSlice, 1, 2, 3)
	fmt.Println("After Append:", intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))
}
