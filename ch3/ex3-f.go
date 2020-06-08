package main

import (
	"fmt"
)

func main() {
	var a, b int8

	fmt.Print("Type A: ")
	fmt.Scan(&a)
	fmt.Print("Type B: ")
	fmt.Scan(&b)

	a, b = b, a //CHANGE THE VALUES

	fmt.Printf("A: %d, B: %d", a, b)
}
