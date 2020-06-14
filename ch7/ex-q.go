package main

import "fmt"

func main() {
	var a [3]string

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%s", &a[i])
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

		var srch string
		fmt.Print("Which name do you want do search? ")
		fmt.Scanf("%s", &srch)

		for begin <= end && find == false {
			mid = (begin + end) / 2

			if srch == a[mid] {
				find = true
			} else {
				if srch < a[mid] {
					end = mid - 1
				} else {
					begin = mid + 1
				}
			}
		}
		if find == true {
			fmt.Printf("\n%s found at [%d]", srch, mid)
		} else {
			fmt.Print("\nCannot found")
		}

		fmt.Print("\nDo you want do search again?(YES/NO) ")
		fmt.Scanf("%s", &answer)
	}
}
