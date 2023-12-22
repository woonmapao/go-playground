package functions

import "fmt"

// Higher-order function that takes a function as a parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Closure example
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Defer and Panic example
func divideWithPanic(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0
		}
	}()
	result = a / b
	return result
}

// Variadic function with deferred function call
func deferredPrint(args ...interface{}) {
	defer fmt.Println("Deferred: End of function")
	fmt.Println("Deferred: Start of function")
	fmt.Println(args...)
}

// Anonymous function example
var square = func(x int) int {
	return x * x
}

func AdvancedFunctionExamples() {
	// Higher-order function example
	sumResult := applyOperation(3, 5, func(a, b int) int {
		return a + b
	})
	fmt.Println("Higher-order function result:", sumResult)

	// Closure example
	counterFunc := counter()
	fmt.Println("Closure count:", counterFunc())
	fmt.Println("Closure count:", counterFunc())

	// Defer and Panic example
	fmt.Println("Divide with Panic result:", divideWithPanic(10, 2))
	fmt.Println("Divide with Panic result:", divideWithPanic(5, 0))

	// Variadic function with deferred function call
	deferredPrint("Hello", 42, 3.14)

	// Anonymous function example
	squared := square(4)
	fmt.Println("Square:", squared)
}
