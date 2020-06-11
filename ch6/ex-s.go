package main

import "fmt"

func main() {
	var a [6]int
	var b [6]int
	var c [12]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A(Only PAIR): ")
		fmt.Scanf("%d", &a[i])

		for (a[i] % 2) != 0 {
			fmt.Print("Type A(Only PAIR): ")
			fmt.Scanf("%d", &a[i])
		}

		fmt.Print("Type B(Only ODD): ")
		fmt.Scanf("%d", &b[i])

		for (b[i] % 3) != 0 {
			fmt.Print("Type B(Only ODD): ")
			fmt.Scanf("%d", &b[i])
		}

		c[i] = a[i]
		c[(i + 6)] = b[i]
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("%dº C: %d\n", (i + 1), c[i])
	}
}
