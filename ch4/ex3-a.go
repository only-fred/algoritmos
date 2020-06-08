package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Type A and B: ")
	fmt.Scanf("%d %d", &a, &b)

	if a < b {
		fmt.Println("A is smaller than B")
	} else {
		fmt.Println("A is bigger than B")
	}
}
