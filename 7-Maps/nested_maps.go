package maps

import "fmt"

// Function to demonstrate nested maps
func NestedMapExample() {
	fmt.Println("Nested Map Example:")

	// Declare and initialize a nested map
	countryPopulations := map[string]map[string]int{
		"USA": {
			"New York":    8336817,
			"Los Angeles": 3980400,
			"Chicago":     2716000,
		},
		"India": {
			"Delhi":     19152334,
			"Mumbai":    12442373,
			"Bangalore": 8443675,
		},
		"China": {
			"Beijing":   21516000,
			"Shanghai":  24183300,
			"Guangzhou": 14043500,
		},
	}

	// Accessing values in a nested map
	populationNY := countryPopulations["USA"]["New York"]
	fmt.Println("Population of New York, USA:", populationNY)

	// Adding a new entry to the nested map
	countryPopulations["Brazil"] = map[string]int{
		"Rio de Janeiro": 6718903,
		"Sao Paulo":      12106920,
	}
	fmt.Println("After Adding Brazil:", countryPopulations)

	// Modifying a value in the nested map
	countryPopulations["India"]["Delhi"] = 20000000
	fmt.Println("After Modifying Delhi, India:", countryPopulations)

	// Deleting an entry from the nested map
	delete(countryPopulations["China"], "Shanghai")
	fmt.Println("After Deleting Shanghai, China:", countryPopulations)
}
