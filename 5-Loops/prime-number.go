package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n < max+1; n++ { // CURRENT ITERATION
		if n == 2 {
			fmt.Println(n)
			continue
			/* 'continue' is used to skip the rest of the CURRENT ITERATION
			in a loop and proceed to the NEXT ITERATION*/
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break // 'break' is used to exit loop prematurely //
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func main() {
	printPrimes(100)
}
