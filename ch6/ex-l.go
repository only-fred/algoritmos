package main

import "fmt"

func main() {
	var a [10]int
	var b [10]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("Type a number: ")
		fmt.Scanf("%d", &a[i])

		b[i] = float64(a[i]) / 2
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("A: %d, B: %g\n", a[i], b[i])
	}
}
