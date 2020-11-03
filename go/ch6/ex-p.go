package main

import "fmt"

func main() {
	var b [12]int
	var a [12]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		if (a[i] % 2) == 0 {
			b[i] = a[i] * 2
		} else {
			b[i] = a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("A: %d -- B: %d\n", a[i], b[i])
	}
}
