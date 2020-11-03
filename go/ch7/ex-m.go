package main

import "fmt"

func main() {
	var a [3]float64
	var b [3]float64
	var c [3]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("A: ")
		fmt.Scanf("%g", &a[i])
		fmt.Print("B: ")
		fmt.Scanf("%g", &b[i])

		c[i] = a[i] + b[i]
	}

	answer := "YES"

	for answer == "yes" || answer == "YES" {
		find := false
		srch := 0.0

		fmt.Print("Which number do you want to find? ")
		fmt.Scanf("%g", &srch)

		for i := 0; i < len(c) && find == false; i++ {
			if srch == c[i] {
				fmt.Printf("\nNumber: %g found at c[%d]", srch, i)

				find = true
			}
		}
		if find == false {
			fmt.Print("\nCannot found your search")
		}

		fmt.Println("\n\nDo you want to search again?(YES/NO)")
		fmt.Scanf("%s", &answer)
	}
}
