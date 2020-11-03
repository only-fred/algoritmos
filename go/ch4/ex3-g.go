package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Print("Type A, B, C and D: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	if (a%2) == 0 && (a%3) == 0 {
		fmt.Printf("%d can divide by 2 and 3\n", a)
	} else {
		fmt.Printf("%d cannot be divided\n", a)
	}

	if (b%2) == 0 && (b%3) == 0 {
		fmt.Printf("%d can divide by 2 and 3\n", b)
	} else {
		fmt.Printf("%d cannot be divided\n", b)
	}

	if (c%2) == 0 && (c%3) == 0 {
		fmt.Printf("%d can divide by 2 and 3\n", c)
	} else {
		fmt.Printf("%d cannot be divided\n", c)
	}

	if (d%2) == 0 && (d%3) == 0 {
		fmt.Printf("%d can divide by 2 and 3\n", d)
	} else {
		fmt.Printf("%d cannot be divided\n", d)
	}
}
