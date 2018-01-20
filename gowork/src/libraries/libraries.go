package main

import (
	"fmt"
	"stringutil"
)

func main() {

	n1, l1 := stringutil.FullName("Sattar", "Hummatli")
	fmt.Printf("Fullname: %v, number of chars: %v\n", n1, l1)

	n2, l2 := stringutil.FullNameNamkedReturn("Ali", "Hummatli")
	fmt.Printf("Fullname: %v, number of chars: %v\n", n2, l2)

}
