package main

import "fmt"

func main() {
	var a [15]int
	var b [15]int
	var c [15]int
	var d [15]int

	for i := 0; i < len(a); i++ {
		fmt.Print("A: ")
		fmt.Scanf("%d", &a[i])

		fmt.Print("B: ")
		fmt.Scanf("%d", &b[i])

		fmt.Print("C: ")
		fmt.Scanf("%d", &c[i])

		d[i] = (a[i] + b[i] + c[i])
	}

	for i := 0; i < (len(a) - 1); i++ {
		for j := i + 1; j < len(a); j++ {
			if d[i] > d[j] {
				d[i], d[j] = d[i], d[j]
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

			if srch == d[mid] {
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
			fmt.Printf("\nNumber found at [%d]", mid)
		} else {
			fmt.Print("\nNumber cannot be found")
		}

		fmt.Print("\nDo you want do search again? (YES/NO)\n")
		fmt.Scanf("%s", &answer)
	}
}
