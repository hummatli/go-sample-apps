package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")

	fmt.Println("Undefed statement")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++

	fmt.Println("Value of x after incrementing it:", x)
}

func myFunc() {
	defer fmt.Println("Defed in the function")
	fmt.Println("Undefed in the function")
}
