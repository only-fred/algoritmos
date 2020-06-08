package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter your number: ")
	fmt.Scanf("%d", &n)

	if n < 0 {
		n = n * (-1)
	}

	fmt.Printf("Your number is: %d", n)
}
