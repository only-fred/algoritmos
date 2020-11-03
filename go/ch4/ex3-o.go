package main

import "fmt"

func main() {
	var n, sum int

	fmt.Print("Number: ")
	fmt.Scanf("%d", &n)

	sum = n * 2

	if sum > 30 {
		fmt.Print("Bigger than 30")
	} else {
		fmt.Print("Smaller than 30")
	}
}
