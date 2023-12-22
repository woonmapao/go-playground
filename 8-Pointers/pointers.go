package variables

import "fmt"

// Function to demonstrate basic usage of pointers
func BasicPointerExample() {
	fmt.Println("Basic Pointer Example:")

	// Declare a variable
	number := 42
	fmt.Println("Original Value:", number)

	// Declare a pointer and assign the address of the variable to it
	var pointerToNumber *int = &number

	// Access the value through the pointer
	fmt.Println("Value through Pointer:", *pointerToNumber)

	// Modify the value through the pointer
	*pointerToNumber = 99
	fmt.Println("Modified Value:", number)
}

// Function to demonstrate pointers with structs
func PointerWithStructExample() {
	fmt.Println("Pointer with Struct Example:")

	// Declare a struct
	type Person struct {
		Name string
		Age  int
	}

	// Declare a variable of type Person
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("Original Person:", person)

	// Declare a pointer to a Person and assign the address of the variable to it
	var pointerToPerson *Person = &person

	// Modify the struct fields through the pointer
	pointerToPerson.Age = 26
	fmt.Println("Modified Person:", person)
}

// Function to demonstrate pointers with slices
func PointerWithSliceExample() {
	fmt.Println("Pointer with Slice Example:")

	// Declare a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", numbers)

	// Declare a pointer to a slice and assign the address of the slice to it
	var pointerToSlice *[]int = &numbers

	// Modify the slice through the pointer
	*pointerToSlice = append(*pointerToSlice, 6, 7, 8)
	fmt.Println("Modified Slice:", numbers)
}
