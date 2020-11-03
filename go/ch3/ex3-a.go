package main

import (
	"fmt"
)

func main() {
	celsius := 0.0

	fmt.Print("Enter Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("The Fahrenheit is: %gº", fahrenheit)
}
