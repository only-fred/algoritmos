package main

import (
	"fmt"
)

func main() {
	var realAmount, dollarAmount, dollarRate float64

	fmt.Print("How much is dolar rate today? ")
	fmt.Scanf("%g", &dollarRate)
	fmt.Print("How much dollars you want do exchange? ")
	fmt.Scanf("%g", &realAmount)

	dollarAmount = realAmount / dollarRate

	fmt.Printf("You exchange R$%g for $%g on rate by %g\n", realAmount, dollarAmount, dollarRate)
}
