package main

import (
	"fmt"
)

func main() {
	fahrenheit := 0.0

	fmt.Print("Enter Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := ((fahrenheit - 32) * 5) / 9

	fmt.Printf("The Celsius is: %gº", celsius)
}
