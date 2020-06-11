package main

import "fmt"

func main() {
	var b [3]float64
	var a [3]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%g", &a[i])

		if (int(a[i]) % 2) == 0 {
			b[i] = a[i] * 2
		} else {
			b[i] = a[i] * 1.5
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("A: %g -- B: %g\n", a[i], b[i])
	}
}
