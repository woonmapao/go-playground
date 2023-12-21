package main

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
type shape interface {
	area() float64
	perimeter() float64
}

// --------------------------------------
type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// --------------------------------------
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 { // (diameter)
	return 2 * 3.14 * c.radius
}

// --------------------------------------

func getAll(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
	fmt.Println("=============")
}

// Type Assertions
func getBetterAll(s shape) (string, float64, float64) {
	c, ok := s.(circle)
	if ok {
		return "circle", c.area(), c.perimeter()
	}
	r, ok := s.(rectangle)
	if ok {
		return "rectangle", r.area(), r.perimeter()
	}
	return "", 0.0, 0.0
}

func main() {

	mysteryShape1 := circle{
		radius: 1.2,
	}

	mysteryShape2 := rectangle{
		width:  2.1,
		height: 1.9,
	}

	getAll(mysteryShape1)

	getAll(mysteryShape2)

	fmt.Println(getBetterAll(mysteryShape1))

	fmt.Println(getBetterAll(mysteryShape2))

}
