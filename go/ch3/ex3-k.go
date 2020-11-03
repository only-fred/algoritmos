package main

import (
	"fmt"
)

func main() {
	var realAmount, dollarAmount, dollarRate float64

	fmt.Print("How much is dolar rate today? ")
	fmt.Scanf("%g", &dollarRate)
	fmt.Print("How much dollars you want do exchange? ")
	fmt.Scanf("%g", &dollarAmount)

	realAmount = dollarAmount * dollarRate

	fmt.Printf("You exchange $%g for R$%g on rate by %g\n", dollarAmount, realAmount, dollarRate)
}
