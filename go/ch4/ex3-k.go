package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	if n >= 1 && n <= 9 {
		fmt.Print("Allow")
	} else {
		fmt.Print("Deny")
	}
}
