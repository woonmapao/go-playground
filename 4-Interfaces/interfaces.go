package interfaces

import "fmt"

/*
1.) keep interfaces small
2.) interfaces should have no

	knowledge of satisfying types
	i.e.
		type car interface {
			Color() string
			Speed() int
			IsSedan() bool		<-----
			IsPickup() bool		<-----
		}
*/
// Define a simple Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a Rectangle struct that implements the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Define a Circle struct that implements the Shape interface
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implement the Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Function to calculate and print the area and perimeter of a shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
