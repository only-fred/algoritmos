package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("A: ")
	fmt.Scanf("%d", &a)
	fmt.Print("B: ")
	fmt.Scanf("%d", &b)

	if b < a {
		a, b = b, a
	}

	fmt.Print("C: ")
	fmt.Scanf("%d", &c)

	if c < b {
		c, b = b, c
	}
	if b < a {
		a, b = b, a
	}

	fmt.Printf("%d %d %d\n", a, b, c)
}
