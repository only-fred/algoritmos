package main

import "fmt"

func main() {
	var a [8]int
	var b [8]int

	for i := 0; i < 8; i++ {
		fmt.Printf("%dº number: ", (i + 1))
		fmt.Scanf("%d", &a[i])
		b[i] = a[i] * 3
	}
	for i := 0; i < 8; i++ {
		fmt.Printf("%dº number: %d\n", (i + 1), a[i])
	}
}
