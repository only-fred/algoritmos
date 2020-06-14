package main

import "fmt"

func main() {
	var a [3]float64
	var b [3]float64
	var c [3]float64
	var d [3]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("A: ")
		fmt.Scanf("%g", &a[i])

		fmt.Print("B: ")
		fmt.Scanf("%g", &b[i])

		fmt.Print("C: ")
		fmt.Scanf("%g", &c[i])

		d[i] = 1

		for j := 0; j < 3; j++ {
			d[i] = d[i] * (a[i] + b[i] + c[i])
		}
	}

	for i := 0; i < (len(d) - 1); i++ {
		for j := i + 1; j < len(d); j++ {
			if d[i] > d[j] {
				d[i], d[j] = d[i], d[j]
			}
		}
	}

	for i := 0; i < len(d); i++ {
		fmt.Printf("[%d]º D: %g\n", i, d[i])
	}
}
