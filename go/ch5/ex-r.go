package main

import "fmt"

func main() {
	var value, bigger, smaller int

	value = 1
	bigger = value
	smaller = value

	for value >= 0 {
		fmt.Print("Number: ")
		fmt.Scanf("%d", &value)

		if value >= 0 {
			if value > bigger {
				value, bigger = bigger, value
			}
			if value < smaller {
				value, smaller = smaller, value
			}
		}
	}

	fmt.Printf("Smaller: %d\nBigger: %d", smaller, bigger)
}
