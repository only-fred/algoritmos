package main

import "fmt"

func main() {
	var celsius [2]float64
	var fahrenheit [2]float64

	for i := 0; i < len(celsius); i++ {
		fmt.Print("Temperature in celsius: ")
		fmt.Scanf("%g", &celsius[i])

		fahrenheit[i] = (celsius[i] * 9 / 5) + 32
	}

	for i := 0; i < len(celsius); i++ {
		fmt.Printf("Celsius: %gº -- Fahrenheit: %gº\n", celsius[i], fahrenheit[i])
	}
}
