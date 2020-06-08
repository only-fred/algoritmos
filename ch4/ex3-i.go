package main

import "fmt"

func main() {
	var a, b, c, d, e, big, small int

	fmt.Print("A: ")
	fmt.Scanf("%d", &a)
	big, small = a, a
	fmt.Print("B: ")
	fmt.Scanf("%d", &b)
	if b > big {
		big = b
	}
	if b < small {
		small = b
	}
	fmt.Print("C: ")
	fmt.Scanf("%d", &c)
	if c > big {
		big = c
	}
	if c < small {
		small = c
	}
	fmt.Print("D: ")
	fmt.Scanf("%d", &d)
	if d > big {
		big = d
	}
	if d < small {
		small = d
	}
	fmt.Print("E: ")
	fmt.Scanf("%d", &e)
	if e > big {
		big = e
	}
	if e < small {
		small = e
	}

	fmt.Printf("Big: %d\nSmall: %d", big, small)
}
