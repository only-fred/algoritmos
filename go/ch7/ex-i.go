package main

import "fmt"

func main() {
	var a [3]int
	var b [3]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		b[i] = a[i] / 2
	}

	for i := 0; i < (len(a) - 1); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("[%d] A: %d\n[%d] B: %d\n", i, a[i], i, b[i])
	}
}
