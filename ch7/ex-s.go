package main

import "fmt"

func main() {
	var a [3]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("A: ")
		fmt.Scanf("%g", &a[i])
	}

	answer := "YES"

	for answer == "YES" || answer == "yes" {
		srch := 0.0
		find := false
		i := 0

		fmt.Print("Which number do you want do search? ")
		fmt.Scanf("%g", &srch)

		for i = 0; i < len(a) && find == false; i++ {
			if srch == a[i] {
				find = true
				break
			}
		}
		if find == true {
			fmt.Printf("Number found at [%d]", i)
		} else {
			fmt.Print("\nCannot found")
		}

		fmt.Print("\nDo you want do search again? (YES/NO)")
		fmt.Scanf("%s", &answer)
	}
}
