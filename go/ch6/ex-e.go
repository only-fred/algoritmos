package main

import "fmt"

func main() {
	var a [15]int
	var b [15]int

	for i := 0; i < len(b); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])
		b[i] = 1

		for j := 0; j < a[i]; j++ {
			b[i] = b[i] * (a[i] - j)
		}
		fmt.Printf("%d fat is: %d\n", a[i], b[i])
	}
}
