package main

import "fmt"

func main() {
	var a [20]string
	var b [30]string
	var c [50]string

	for i := 0; i < len(a); i++ {
		fmt.Print("Type name to first array(A): ")
		fmt.Scanf("%s", &a[i])

		c[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		fmt.Print("Type name to second array(B): ")
		fmt.Scanf("%s", &b[i])

		c[(i + 2)] = b[i]
	}

	for i := 0; i < (len(c) - 1); i++ {
		for j := i + 1; j < len(c); j++ {
			if c[i] < c[j] {
				c[i], c[j] = c[j], c[i]
			}
		}
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("[%d]º C: %s\n", i, c[i])
	}
}
