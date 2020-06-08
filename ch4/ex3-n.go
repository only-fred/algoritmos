package main

import "fmt"

func main() {
	var a, b, c, sum int

	fmt.Print("A, B and C: ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	sum = a + b + c

	if sum >= 100 {
		fmt.Printf("Bigger or equals 100")
	} else {
		fmt.Print("Smaller than 100")
	}
}
