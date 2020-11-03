package main

import "fmt"

func main() {
	var a [10]int
	var b [10]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		for j := 0; j > a[i]; j++ {
			fmt.Print("A positive number, idiot: ")
			fmt.Scanf("%d", &a[i])

			if a[i] > 0 {
				break
			}
		}
	}
	for i := 0; i < len(b); i++ {
		b[i] = a[i] - a[i] - a[i]
		fmt.Print(b[i], "\n")
	}
}
