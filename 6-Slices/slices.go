package main

import "fmt"

func main() {
	// WE USE SLICES > ARRAY MOST OF THE TIME ! (99%)

	// arrays
	var myArray [10]int // ---> [0 0 0 0 0 0 0 0 0 0]
	primes := [6]int{2, 3, 5, 7, 11, 13}

	myArray[1] = 1
	primes[1] = 1

	// slices
	var mySlice []int
	fmt.Println(mySlice)

	// func make([]Type, len, cap) []T
	mySlice2 := make([]int, 5, 10)
	// the capacity argument is usually omitted
	// and defaults to the length
	fmt.Println(len(mySlice2))

	matrix1 := createMatrix(5, 5)
	fmt.Println(matrix1)
	matrix2 := createMatrix2(5, 5)
	fmt.Println(matrix2)

	// don't do this !
	// someSlice = append(otherSlide, element)

}

func createMatrix(rows, cols int) [][]int {

	myMatrix := make([][]int, rows)
	for i := 0; i < cols; i++ {
		myMatrix[i] = make([]int, cols)
	}
	return myMatrix
}

func createMatrix2(rows, cols int) [][]int {

	myMatrix := make([][]int, rows)

	for i := 0; i < rows; i++ {

		myMatrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			myMatrix[i][j] = i * j
		}
	}
	return myMatrix
}
