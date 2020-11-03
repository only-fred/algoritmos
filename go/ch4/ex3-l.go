package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	if !(n > 3) {
		fmt.Printf("%d", n)
	}
}
