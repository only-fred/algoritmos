package main

import "fmt"

func main() {
	var a [15]int

	for i := 0; i < len(a); i++ {
		fmt.Print("A: ")
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < (len(a) - 1); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	answer := "YES"

	for answer == "YES" || answer == "yes" {
		begin := 0
		mid := 0
		end := len(a)
		find := false
		srch := 0

		fmt.Print("Which number do you want do search? ")
		fmt.Scanf("%d", &srch)

		for begin <= end && find == false {
			mid = (begin + end) / 2

			if srch == a[mid] {
				find = true
			} else {
				if srch > a[mid] {
					end = mid - 1
				} else {
					begin = mid + 1
				}
			}
		}
		if find == true {
			fmt.Printf("\n%d found at a[%d]", srch, mid)
		} else {
			fmt.Print("\nCannot found")
		}

		fmt.Print("\nDo you want do search again? (YES/NO)\n")
		fmt.Scanf("%s", &answer)
	}
}
