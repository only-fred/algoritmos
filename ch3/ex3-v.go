package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = 3.14159

	var radius float64

	fmt.Print("Type Radius: ")
	fmt.Scanf("%g", &radius)

	volume := (4 / 3) * PI * (math.Pow(radius, 3))

	fmt.Printf("Volume is: %g³", volume)
}
