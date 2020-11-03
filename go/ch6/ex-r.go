package main

import "fmt"

func main() {
	var a [6]int
	var b [6]int
	var c [12]int
	var d [12]int

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])
		fmt.Print("Type B: ")
		fmt.Scanf("%d", &b[i])

		if (a[i] % 2) == 0 {
			c[i] = a[i]
		} else {
			d[i] = a[i]
		}

		if (b[i] % 2) == 0 {
			c[i] = b[i]
		} else {
			d[i] = b[i]
		}
	}

	for i := 0; i < len(c); i++ {
		if c[i] == 0 && d[i] == 0 {
			break
		}
		fmt.Printf("C: %d -- D: %d\n", c[i], d[i])
	}
}
