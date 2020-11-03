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

		sum := a[i] + b[i]

		c[i] = sum * sum
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("%dº C: %d\n", (i + 1), c[i])
	}
}
