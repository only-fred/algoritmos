package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {

		if (i % 3) == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
