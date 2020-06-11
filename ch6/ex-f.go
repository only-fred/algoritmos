package main

import "fmt"

func main() {
	var a [5]int
	var b [5]int
	var c [10]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])
		fmt.Print("Type B: ")
		fmt.Scanf("%d", &b[i])

		c[i] = a[i]
		c[(i + 5)] = b[i]
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("%d\n", c[i])
	}
}
