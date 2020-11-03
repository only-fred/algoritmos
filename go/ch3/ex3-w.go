package main

import "fmt"

func main() {
	var feet, meters float64

	fmt.Print("How many feet? ")
	fmt.Scanf("%g", &feet)

	meters = (feet * 0.3048)

	fmt.Printf("Meters: %gm", meters)
}
