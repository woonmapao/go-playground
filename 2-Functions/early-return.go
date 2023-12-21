package main

import (
	"errors"
)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
		// early return, can clean up code
	}
	return dividend / divisor, nil
}
