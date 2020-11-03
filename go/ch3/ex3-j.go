package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Type A: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Type B: ")
	fmt.Scanf("%f", &b)

	a = math.Pow(a, 2)
	b = math.Pow(b, 2)

	if a > b {
		fmt.Printf("A(%g) its bigger than B(%g)", a, b)
	} else {
		fmt.Printf("A(%g) its smaller than B(%g)", a, b)
	}
}
