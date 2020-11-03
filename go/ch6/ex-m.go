package main

import "fmt"

func main() {
	var tab [10]int
	a := 0

	fmt.Print("Type a number (1-10): ")
	fmt.Scanf("%d", &a)

	for i := 0; i < len(tab); i++ {
		tab[i] = (i + 1) * a

		fmt.Printf("%d x %d = %d\n", a, i, tab[i])
	}
}
