package main

import "fmt"

func main() {
	var a [20]string
	var b [30]string
	var c [50]string

	for i := 0; i < len(a); i++ {
		fmt.Print("Type a name: ")
		fmt.Scanf("%s", &a[i])

		c[i] = a[i]

	}
	for i := 0; i < len(b); i++ {
		fmt.Print("Type a name: ")
		fmt.Scanf("%s", &b[i])

		c[(i + 2)] = b[i]
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("%s\n", c[i])
	}
}
