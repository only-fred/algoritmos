package main

import (
	"fmt"
	"math"
)

func main() {
	var i float64

	fmt.Print("Enter the value: ")
	fmt.Scanf("%g", &i)

	i = math.Pow(i, 2)

	fmt.Printf("The value squared is: %g", i)
}
