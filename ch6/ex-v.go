package main

import "fmt"

func main() {
	var a [30]int
	var pair, odd int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		if (a[i] % 2) == 0 {
			pair = pair + 1
		} else {
			odd = odd + 1
		}
	}

	fmt.Printf("%d odds\n%d pairs", odd, pair)
}
