package main

import (
	"fmt"
)

func main() {
	var x float64 = 42
	var result string

	if x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)

	if y := -42; y < 0 {
		result = "Less than zero"
	} else if y == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)
}
