package functions

import "fmt"

// Function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divideAndRemainder(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// Function with a variadic parameter
func sum(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

// Function with a named return value
func multiply(a, b int) (product int) {
	product = a * b
	return
}

func BasicFunctionExamples() {
	// Call functions and print results
	fmt.Println("Addition:", add(5, 3))

	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Printf("Division: Quotient - %d, Remainder - %d\n", quotient, remainder)

	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	product := multiply(4, 6)
	fmt.Println("Multiplication:", product)
}
