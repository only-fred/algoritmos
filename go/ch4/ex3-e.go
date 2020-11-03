package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Printf("Type A: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Type B: ")
	fmt.Scanf("%f", &b)
	fmt.Printf("Type C: ")
	fmt.Scanf("%f", &c)

	delta := b*b - 4*a*c

	x1 := (-b + math.Sqrt(delta)) / (2 * a)
	x2 := (-b - math.Sqrt(delta)) / (2 * a)

	if delta < 0 {
		fmt.Printf("The equation doesn't have square root")
	} else {
		fmt.Printf("x1: %f", x1)
		fmt.Printf("x2, %f", x2)
	}
}
