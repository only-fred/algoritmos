package main

import "fmt"

func main() {
	var a [20]int
	var b [20]int
	var c [20]int

	for i := 0; i < len(c); i++ {
		fmt.Printf("Type %dº A:", (i + 1))
		fmt.Scanf("%d", &a[i])
		fmt.Printf("Type %dº B:", (i + 1))
		fmt.Scanf("%d", &b[i])
	}
	for i := 0; i < len(c); i++ {
		c[i] = a[i] - b[i]

		fmt.Printf("%dº C is: %d\n", (i + 1), c[i])
	}
}
