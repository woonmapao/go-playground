package variables

import "fmt"

func BasicVariableExamples() {
	// Declare and initialize variables
	var a int = 5
	var b float64 = 3.14
	var c string = "Hello, Go!"

	// Short declaration syntax
	x := 10
	y := 2.71
	z := "Short declaration syntax"

	// Print the values
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Update variable values
	a = 7
	x = 20

	// Print updated values
	fmt.Println("Updated a:", a)
	fmt.Println("Updated x:", x)
}
