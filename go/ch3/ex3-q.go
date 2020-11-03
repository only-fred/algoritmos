package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = 3.1415265

	var radius float64

	fmt.Print("Type the Radius: ")
	fmt.Scanf("%g", &radius)

	area := PI * math.Pow(radius, 2)

	fmt.Printf("The area is: %g³", area)
}
