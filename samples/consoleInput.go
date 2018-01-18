package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Simple input scan
	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)

	//Read by Reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

}
