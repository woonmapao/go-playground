package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 5
	z := &x
	myString := []string{"hello", "oh", "shoot"}

	fmt.Println(z)
	for _, message := range myString {

		removeProfanity(&message)
		fmt.Println(message)
	}

}

func removeProfanity(message *string) {
	if message == nil {
		return
	}

	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}
