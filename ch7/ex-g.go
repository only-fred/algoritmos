package main

import "fmt"

func main() {
	var a [20]int
	var b [20]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		b[i] = a[i] + 2
	}

	answer := "YES"

	for answer == "YES" || answer == "yes" {
		srch := 0
		begin := 0
		mid := 0
		end := len(a)
		find := false

		fmt.Print("Which number do you like do search? ")
		fmt.Scanf("%d", &srch)

		for begin <= end && find == false {
			mid = (begin + end) / 2

			if srch == b[mid] {
				find = true
			} else {
				if srch < b[mid] {
					end = mid - 1
				} else {
					begin = mid + 1
				}
			}
		}
		if find == true {
			fmt.Printf("Number find on index [%d]\n\n", mid)
		} else {
			fmt.Print("Number doesn't found\n\n")
		}
		fmt.Print("Want search again (YES/NO)?\n")
		fmt.Scanf("%s", &answer)
	}
}
