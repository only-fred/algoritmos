package main

import "fmt"

func main() {
	var a [20]int
	var b [20]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type a number: ")
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < len(b); i++ {
		b[i] = a[len(a)-i-1]
		fmt.Printf("B: %d\n", b[i])
	}
}
