package main

import "fmt"

func main() {
	sum, result := 0, 0

	for i := 1; i <= 500; i++ {
		sum = sum + 1

		if (sum % 2) == 0 {
			result = result + sum
		}
	}
	fmt.Printf("%d\n", result)
}
