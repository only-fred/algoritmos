package main

import (
	"fmt"
)

func main() {
	var travelTime, mediumVelocity, distance float64

	fmt.Print("How long will be your travel? ")
	fmt.Scan(&travelTime)
	fmt.Print("What will be your medium velocity? ")
	fmt.Scan(&mediumVelocity)

	distance = travelTime * mediumVelocity

	var usedLiters float64 = distance / 12

	fmt.Printf("You will use: %f liters", usedLiters)
}
