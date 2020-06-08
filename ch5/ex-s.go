package main

import "fmt"

func main() {
	var dividend, divider, result, rest int

	dividend = 9
	divider = 3

	for i := 0; i < 3; i++ {
		fmt.Printf("Dividend: %d\n", dividend)
		fmt.Printf("Divider: %d\n", divider)
		fmt.Printf("Result: %d\n", result)
		fmt.Printf("Rest: %d\n\n", rest)

		dividend = dividend / divider

		result = dividend / divider
	}
}
