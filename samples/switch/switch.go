package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
	case 1:
		result = "It's Sunday!"
	case 7:
		result = "It's Saturday!"
	default:
		result = "It's a weekday!"
	}

	fmt.Println("Day", dow, ",", result)

	//switch with variable init

	result = ""

	switch dow2 := rand.Intn(6) + 1; dow2 {
	case 1:
		result = "2, It's Sunday!"
	case 7:
		result = "2, It's Saturday!"
	default:
		result = "2, It's a weekday!"
	}

	fmt.Println(result)

	x := -42
	switch {
	case x < 42:
		result = "Less than zero"
	case x == 0:
		result = "Equal to zero"
	default:
		result = "Greater than zero"
	}
	fmt.Println(result)
}
