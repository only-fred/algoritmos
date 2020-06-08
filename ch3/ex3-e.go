package main

import (
	"fmt"
)

func main() {
	var value, rate, time float64

	fmt.Print("What's the value? ")
	fmt.Scan(&value)
	fmt.Print("What's the rate? ")
	fmt.Scan(&rate)
	fmt.Print("How long you'll pay (in days)? ")
	fmt.Scan(&time)

	installment := value + (value * (rate / 100) * time)

	rate = installment - value

	fmt.Printf("The total installment is: %f\nThe rate is: %f", installment, rate)
}
