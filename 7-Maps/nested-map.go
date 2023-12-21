package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {

	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		firstChar := rune(name[0]) // first char as rune

		// Initialize the inner map if it doesn't exist
		if _, ok := nameCounts[firstChar]; !ok {
			nameCounts[firstChar] = make(map[string]int)
		}
		nameCounts[firstChar][name]++
	}
	return nameCounts

}

func main() {
	// Mock some data
	names := []string{"Billy", "Bob", "Billy", "Joe"}

	// Use the function to get name counts
	result := getNameCounts(names)

	// Print the result
	fmt.Println(result)
}
