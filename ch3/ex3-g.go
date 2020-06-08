package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, sum, mult int32

	fmt.Print("Enter A value: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter B value: ")
	fmt.Scanf("%d", &b)
	fmt.Print("Enter C value: ")
	fmt.Scanf("%d", &c)
	fmt.Print("Enter D value: ")
	fmt.Scanf("%d", &d)

	sum = sum + (a + b) + (a + c) + (a + d)
	sum = sum + (b + c) + (b + d)
	sum = sum + (c + d)

	mult = (a * b) * (a * c) * (a * d)
	mult = mult * (b * c) * (b * d)
	mult = mult * (c * d)

	fmt.Printf("Sum: %d, Mult: %d", sum, mult)
}
