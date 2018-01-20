package main

import "fmt"

func main() {
	doSomething()

	sum := addValues(23, 45)
	fmt.Println("Sum:", sum)

	sum = addAllValues(1, 2, 3)
	fmt.Println("Sum All:", sum)
}

func doSomething() {
	fmt.Println("Doing something!")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addValues2(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}
