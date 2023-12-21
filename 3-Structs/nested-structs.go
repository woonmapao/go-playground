package main

import "fmt"

type car struct {
	Model      string
	Height     int
	Width      int
	FrontWheel wheel
	BackWheel  wheel
}

type wheel struct {
	Radius   int
	Material string
}

func main() {

	myCar := car{}
	myCar.FrontWheel.Radius = 5

	myCar2 := car{
		Model:  "camry",
		Height: 177,
		Width:  121,
		FrontWheel: wheel{
			Radius:   10,
			Material: "Carbon",
		},
		BackWheel: wheel{
			Radius:   11,
			Material: "Graphite",
		},
	}

	fmt.Println(myCar2)
}
