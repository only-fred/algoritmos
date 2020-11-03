package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Print("Type A:")
	fmt.Scanf("%d", &a)
	fmt.Print("Type B:")
	fmt.Scanf("%d", &b)
	fmt.Print("Type C:")
	fmt.Scanf("%d", &c)
	fmt.Print("Type D:")
	fmt.Scanf("%d", &d)

	p := a + c
	s := b + d

	fmt.Printf("The sum of A and C is: %d, the sum of B and D is: %d", p, s)
}
