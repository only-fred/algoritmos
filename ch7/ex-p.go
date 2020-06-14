package main

import "fmt"

func main() {
	var a [3]string
	var b [3]string

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A(string): ")
		fmt.Scanf("%s", &a[i])

		b[(2 - i)] = a[i]
	}

	for i := 0; i < (len(b) - 1); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}

	for i := 0; i < len(b); i++ {
		fmt.Printf("B[%d]: %s\n", i, b[i])
	}
}
