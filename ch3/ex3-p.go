package main

import "fmt"

func main() {
	var mountPayment, readjustment, newPayment float64

	fmt.Print("Type your mount payment: $")
	fmt.Scanf("%g", &mountPayment)
	fmt.Print("Type the readjustment: ")
	fmt.Scanf("%g", &readjustment)

	newPayment = mountPayment + (mountPayment * (readjustment / 100))

	fmt.Printf("Your new payment is: $%g", newPayment)
}
