package main

import "fmt"

func main() {
	var a [2]int
	var b [2]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type number: ")
		fmt.Scanf("%d", &a[i])

		for j := 0; j <= a[i]; j++ {
			b[i] = b[i] + j
		}
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("B: %d\n", b[i])
	}
}
