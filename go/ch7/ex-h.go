package main

import "fmt"

func main() {
	var a [2]int
	var b [2]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type a negative number to A: ")
		fmt.Scanf("%d", &a[i])

		for a[i] >= 0 {
			fmt.Print("Try again, a negative number to A: ")
			fmt.Scanf("%d", &a[i])
		}
		b[i] = a[i] * (-1)
	}

	for i := 0; i < (len(b) - 1); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] < b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}

	for i := 0; i < len(b); i++ {
		fmt.Printf("[%d] B: %d", i, b[i])
	}
}
