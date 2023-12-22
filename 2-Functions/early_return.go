package functions

import (
	"errors"
	"fmt"
)

// Function with early return for validation
func validateInput(input int) error {
	if input < 0 {
		return errors.New("input must be non-negative")
	}
	return nil
}

// Function with early return and named return value for readability
func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result = a / b
	return
}

// Function with multiple early returns
func processNumbers(numbers []int) (result int, err error) {
	if len(numbers) == 0 {
		err = errors.New("empty input slice")
		return
	}

	for _, num := range numbers {
		if num < 0 {
			err = errors.New("negative number found")
			return
		}
		result += num
	}

	return
}

func EarlyReturnExamples() {
	// Early return for validation
	err := validateInput(-5)
	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	// Early return with named return value
	quotient, err := divide(8, 2)
	if err != nil {
		fmt.Println("Division failed:", err)
		return
	}
	fmt.Println("Division result:", quotient)

	// Multiple early returns
	numbers := []int{1, 2, 3, 4}
	sum, err := processNumbers(numbers)
	if err != nil {
		fmt.Println("Processing numbers failed:", err)
		return
	}
	fmt.Println("Sum of numbers:", sum)
}
