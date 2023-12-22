package maps

import "fmt"

// Function to demonstrate basic operations on maps
func BasicMapOperations() {
	fmt.Println("Basic Map Operations Example:")

	// Declare and initialize a map
	studentGrades := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}

	// Accessing values in a map
	aliceGrade := studentGrades["Alice"]
	fmt.Println("Alice's Grade:", aliceGrade)

	// Adding a new entry to the map
	studentGrades["David"] = 88
	fmt.Println("After Adding David:", studentGrades)

	// Modifying a value in the map
	studentGrades["Bob"] = 88
	fmt.Println("After Modifying Bob's Grade:", studentGrades)

	// Deleting an entry from the map
	delete(studentGrades, "Charlie")
	fmt.Println("After Deleting Charlie:", studentGrades)
}

// Function to demonstrate iterating over a map
func IterateOverMap() {
	fmt.Println("Iterate Over Map Example:")

	// Declare and initialize a map
	countryCapitals := map[string]string{
		"USA":    "Washington, D.C.",
		"India":  "New Delhi",
		"France": "Paris",
	}

	// Iterate over the map using range
	for country, capital := range countryCapitals {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}
}

// Function to check if a key exists in a map
func CheckKeyExistence() {
	fmt.Println("Check Key Existence Example:")

	// Declare and initialize a map
	languages := map[string]string{
		"Go":     "Awesome",
		"Python": "Cool",
		"Java":   "Sturdy",
	}

	// Check if a key exists
	if value, exists := languages["Go"]; exists {
		fmt.Println("Go is", value)
	} else {
		fmt.Println("Go not found")
	}

	// Check if a key doesn't exist
	if value, exists := languages["JavaScript"]; exists {
		fmt.Println("JavaScript is", value)
	} else {
		fmt.Println("JavaScript not found")
	}
}
