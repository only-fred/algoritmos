package main

import "fmt"

func main() {
	var a [12]int

	for i := 0; i < len(a); i++ {
		fmt.Print("A: ")
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < (len(a)); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				x := a[i]
				a[i] = a[j]
				a[j] = x
			}
		}
		fmt.Printf("%d\n", a[i])
	}
}
