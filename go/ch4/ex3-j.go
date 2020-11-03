package main

import "fmt"

func main() {
	var n int

	fmt.Print("Number: ")
	fmt.Scanf("%d", &n)

	if (n % 2) == 0 {
		fmt.Print("Even")
	} else {
		fmt.Print("Odd")
	}
}
