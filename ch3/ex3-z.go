package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Type A: ")
	fmt.Scanf("%g", &a)
	fmt.Print("Type B: ")
	fmt.Scanf("%g", &b)

	result := math.Pow((a / 2), 2)

	fmt.Printf("Result: %g", result)
}
