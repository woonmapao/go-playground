package structs

import "fmt"

// Address struct representing a person's address
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// Contact struct representing a person's contact information
type Contact struct {
	Email   string
	Phone   string
	Address Address // Embedding the Address struct
}

// Person struct with nested structs
type PersonWithContact struct {
	FirstName string
	LastName  string
	Age       int
	Contact   Contact // Embedding the Contact struct
}

// Function to create and print a person with contact information
func CreateAndPrintPersonWithContact() {
	// Create a person with nested structs
	person := PersonWithContact{
		FirstName: "Patarapong",
		LastName:  "Limvipaveeanan",
		Age:       25,
		Contact: Contact{
			Email: "patarapong@example.com",
			Phone: "555-1234",
			Address: Address{
				Street:  "123 Main St",
				City:    "Anytown",
				State:   "CAT",
				ZipCode: "12345",
			},
		},
	}

	// Print the person details with contact information
	fmt.Println("Person with Contact Details:")
	fmt.Println("First Name:", person.FirstName)
	fmt.Println("Last Name:", person.LastName)
	fmt.Println("Age:", person.Age)
	fmt.Println("Email:", person.Contact.Email)
	fmt.Println("Phone:", person.Contact.Phone)
	fmt.Println("Address:")
	fmt.Println("  Street:", person.Contact.Address.Street)
	fmt.Println("  City:", person.Contact.Address.City)
	fmt.Println("  State:", person.Contact.Address.State)
	fmt.Println("  Zip Code:", person.Contact.Address.ZipCode)
}
