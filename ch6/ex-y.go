package main

import "fmt"

func main() {
	var a [15]int
	var pair int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		if (a[i] % 2) == 0 {
			pair = pair + 1
		}
	}

	fmt.Printf("%d pairs", pair)
}
