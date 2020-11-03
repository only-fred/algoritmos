package main

import "fmt"

func main() {
	var a [10]int
	var b [10]int
	var c [20]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A (Only if divisible by 2 and 3): ")
		fmt.Scanf("%d", &a[i])
		for (a[i]%2) != 0 || (a[i]%3) != 0 {
			fmt.Print("Type A (Only if divisible by 2 and 3): ")
			fmt.Scanf("%d", &a[i])
		}

		fmt.Print("Type B (Only if divisible by 5): ")
		fmt.Scanf("%d", &b[i])
		for (b[i] % 5) != 0 {
			fmt.Print("Type B (Only if divisible by 5): ")
			fmt.Scanf("%d", &b[i])
		}

		c[i] = a[i]
		c[(i + 1)] = b[i]
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("%dº C: %d\n", (i + 1), c[i])
	}
}
