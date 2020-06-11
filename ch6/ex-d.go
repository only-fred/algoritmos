package main

import "fmt"

func main() {
	var a [15]int
	var b [15]int

	for i := 0; i < len(b); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		b[i] = a[i] * a[i]

		fmt.Printf("%dº A squared is: %d\n", (i + 1), b[i])
	}
}
