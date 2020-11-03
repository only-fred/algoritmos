package main

import "fmt"

func main() {
	var a [15]int
	var b [15]int
	var c [15]int
	var d [15]int
	var e [15]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		fmt.Print("Type B: ")
		fmt.Scanf("%d", &b[i])

		c[i] = c[i] + a[i]

		d[i] = 1
		for j := 0; j < b[i]; j++ {
			d[i] = d[i] * (b[i] - j)
		}

		e[i] = (d[i] - c[i]) + (a[i] + b[i])
	}

	for i := 0; i < (len(e) - 1); i++ {
		for j := i + 1; j < len(e); j++ {
			if e[i] > e[j] {
				e[i], e[j] = e[j], e[i]
			}
		}
	}

	for i := 0; i < len(e); i++ {
		fmt.Printf("[%d] E: %d\n", i, e[i])
	}
}
