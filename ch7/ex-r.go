package main

import "fmt"

func main() {
	var a [10]string
	var b [5]string
	var c [15]string

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%s", &a[i])

		if i < len(b) {
			fmt.Print("Type B: ")
			fmt.Scanf("%s", &b[i])
		}
	}

	for i := 0; i < len(a); i++ {
		c[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		c[(10 + i)] = b[i]
	}

	for i := 0; i < (len(c) - 1); i++ {
		for j := i + 1; j < len(c); j++ {
			if c[i] > c[j] {
				c[i], c[j] = c[j], c[i]
			}
		}
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("[%d] C: %s\n", i, c[i])
	}
}
