package main

import "fmt"

func main() {
	var n1, n2, n3, n4, ne float64

	fmt.Print("Enter N1, N2, N3, N4: ")
	fmt.Scanf("%g %g %g %g", &n1, &n2, &n3, &n4)

	md1 := (n1 + n2 + n3 + n4) / 4

	if md1 >= 7 {
		fmt.Printf("Aproved, Media: %g", md1)
	} else {
		fmt.Print("Enter your 5th test: ")
		fmt.Scanf("%g", &ne)
		md2 := (md1 + ne) / 2

		if md2 >= 5 {
			fmt.Printf("Aproved in exam, media: %g", md2)
		} else {
			fmt.Printf("Reproved, media: %g", md2)
		}
	}
}
