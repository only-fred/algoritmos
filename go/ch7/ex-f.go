package main

import (
	"fmt"
)

func main() {
	var a [30]float64
	var b [30]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%g", &a[i])

		b[i] = 1

		for j := 0; j < 3; j++ {
			b[i] = b[i] * a[i]
		}
	}

	answer := "YES"

	for answer == "YES" || answer == "yes" {
		find := false
		srch := 0.0
		resultIndex := 0

		fmt.Print("Which number do you like do find? ")
		fmt.Scanf("%g", &srch)

		for i := 0; i < len(b) && find == false; i++ {
			if srch == b[i] {
				find = true

				resultIndex = i
			}
		}

		if find == true {
			fmt.Printf("%g found at [%d]\n", b[resultIndex], resultIndex)
		} else {
			fmt.Print("Your search cannot be found\n")
		}

		fmt.Print("Do you like to search again? ")
		fmt.Scanf("%s", &answer)
	}
}
