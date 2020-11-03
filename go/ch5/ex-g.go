package main

import (
	"fmt"
)

func main() {
	expoent := 0
	test := 1

	for i := 0; i <= 15; i++ {
		expoent = test
		test = expoent * 3
		result := expoent
		fmt.Print(result, "\n")
	}
}
