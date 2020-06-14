package main

import "fmt"

func main() {
	var a [15]int
	var b [15]int

	for i := 0; i < len(a); i++ {
		fmt.Print("A: ")
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < len(b); i++ {
		b[i] = 1
		for j := 0; j < a[i]; j++ {
			b[i] = b[i] * (a[i] - j)
		}
	}

	for i := 0; i < (len(b) - 1); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				x := b[i]
				b[i] = b[j]
				b[j] = x
			}
		}
		fmt.Printf("[%dº] -> %d\n", i, b[i])
	}
}
