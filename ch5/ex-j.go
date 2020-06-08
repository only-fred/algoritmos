package main

import "fmt"

func main() {
	var celsius, fahrenheit [11]float64

	for i := 1; i < 11; i++ {
		celsius[i] = float64(i) * 10

		fahrenheit[i] = (celsius[i]*9 + 160) / 5
	}

	for i := 1; i < 11; i++ {
		fmt.Printf("Celsius: %fº, Fahrenheit: %fº\n", celsius[i], fahrenheit[i])
	}
}
