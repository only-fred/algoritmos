package main

import "fmt"

func main() {
	var n1, n2, n3, n4 float64

	fmt.Print("Enter N1, N2, N3, N4: ")
	fmt.Scanf("%g %g %g %g", &n1, &n2, &n3, &n4)

	md := (n1 + n2 + n3 + n4) / 4

	if md >= 5 {
		fmt.Print("Aproved ")
	} else {
		fmt.Print("Reproved ")
	}
	fmt.Printf("%g", md)
}
