package variables

import "fmt"

// Custom type declarations
type Celsius float64
type Fahrenheit float64

func TemperatureConversion() {
	// Using custom types
	var freezingC Celsius = 0
	var boilingC Celsius = 100

	// Convert Celsius to Fahrenheit
	freezingF := celsiusToFahrenheit(freezingC)
	boilingF := celsiusToFahrenheit(boilingC)

	// Print the results
	fmt.Printf("Freezing point in Fahrenheit: %v\n", freezingF)
	fmt.Printf("Boiling point in Fahrenheit: %v\n", boilingF)
}

// Helper function to convert Celsius to Fahrenheit
func celsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
