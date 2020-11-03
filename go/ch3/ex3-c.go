package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = 3.14159
	var volume, radius, height float64

	fmt.Print("Enter the radius of can: ")
	fmt.Scan(&radius)
	fmt.Print("Enter the height of can: ")
	fmt.Scan(&height)

	volume = (PI * math.Pow(radius, 2) * height)

	fmt.Printf("The volume of the can is: %f", volume)
}
