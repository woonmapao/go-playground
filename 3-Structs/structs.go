package structs

import "fmt"

// Define a simple person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Function to create and print a person
func CreateAndPrintPerson() {
	// Create a person using the struct
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Print the person details
	fmt.Println("Person Details:")
	fmt.Println("First Name:", person.FirstName)
	fmt.Println("Last Name:", person.LastName)
	fmt.Println("Age:", person.Age)
}
