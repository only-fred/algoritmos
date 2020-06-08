package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Type A: ")
	fmt.Scanf("%g", &a)
	fmt.Print("Type B: ")
	fmt.Scanf("%g", &b)

	r1 := a + b
	r2 := a - b
	r3 := a * b
	r4 := a / b

	fmt.Printf("Sum: %g\nSub: %g\nMult: %g\nDiv: %g", r1, r2, r3, r4)
}
