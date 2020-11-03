package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Type A: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Type B: ")
	fmt.Scanf("%f", &b)
	fmt.Print("Type C: ")
	fmt.Scanf("%f", &c)

	a = math.Pow(a, 2)
	b = math.Pow(b, 2)
	c = math.Pow(c, 2)

	result := a + b + c

	result = math.Pow(result, 2)

	fmt.Printf("The result is: %g", result)
}
