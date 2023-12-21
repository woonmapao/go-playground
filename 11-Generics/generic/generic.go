package main

func main() {

	// Generics will probably be used more heavily in LIBRARY PACKAGES

}

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func getLast[T any](slice []T) T {
	if len(slice) == 0 {
		var myZero T
		return myZero
	}
	return slice[len(slice)-1]
}
