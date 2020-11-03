package main

import "fmt"

func main() {
	for i := 1; i <= 200; i++ {
		if (i % 4) == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
