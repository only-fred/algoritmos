package main

import "fmt"

func main() {
	var highestValue, lowestValue, sumMedia, media float64
	var celsius [3]float64

	for i := 0; i < len(celsius); i++ {
		fmt.Print("Temperature in celsius: ")
		fmt.Scanf("%g", &celsius[i])

		lowestValue = celsius[i]

		sumMedia = sumMedia + celsius[i]
	}

	for i := 0; i < len(celsius); i++ {
		if highestValue <= celsius[i] {
			highestValue = celsius[i]
		}
		if lowestValue >= celsius[i] {
			lowestValue = celsius[i]
		}
	}
	media = sumMedia / 3
	fmt.Printf("Highest value: %gº\nLowest value: %gº\nMedia: %g", highestValue, lowestValue, media)
}
