package main

import "fmt"

func main() {
	var a [3]float64
	var b [3]float64
	var c [3]float64

	for i := 0; i < len(a); i++ { //READ A
		fmt.Print("A: ")
		fmt.Scanf("%g", &a[i])
	}
	for i := 0; i < (len(a) - 1); i++ { //ORDER A
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	for i := 0; i < len(b); i++ { //READ B
		fmt.Print("B: ")
		fmt.Scanf("%g", &b[i])
	}
	for i := 0; i < (len(b) - 1); i++ { //ORDER B
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}

	for i := 0; i < len(c); i++ {
		c[i] = a[i] + b[i]
	}

	for i := 0; i < (len(c) - 1); i++ {
		for j := i + 1; j < len(c); j++ {
			if c[i] < c[j] {
				c[i], c[j] = c[j], c[i]
			}
		}
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("[%d]º C: %g\n", i, c[i])

	}
}
