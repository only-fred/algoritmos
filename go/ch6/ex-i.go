package main

import "fmt"

func main() {
	var a [5]float64
	var b [5]float64
	var c [5]float64
	var d [15]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("Type number: ")
		fmt.Scanf("%g", &a[i])

		d[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		fmt.Print("Type number: ")
		fmt.Scanf("%g", &b[i])

		d[(i + len(a))] = b[i]
	}
	for i := 0; i < len(c); i++ {
		fmt.Print("Type number: ")
		fmt.Scanf("%g", &c[i])
		d[(i + len(a) + len(b))] = c[i]
	}

	for i := 0; i < len(d); i++ {
		fmt.Printf("%g\n", d[i])
	}
}
