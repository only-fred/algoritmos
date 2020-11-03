package main

import "fmt"

func main() {
	var a [10]int
	var b [10]int
	var c [10]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])
		fmt.Print("Type B: ")
		fmt.Scanf("%d", &b[i])

		c[i] = (a[i] * a[i]) + (b[i] * b[i])
	}

	for i := 0; i < (len(c) - 1); i++ {
		for j := i + 1; j < len(c); j++ {
			if c[i] < c[j] {
				c[i], c[j] = c[j], c[i]
			}
		}
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("[%d] C: %d\n", i, c[i])
	}
}
