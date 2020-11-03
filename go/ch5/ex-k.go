package main

import "fmt"

func main() {
	var square [64]int
	var sum int
	for i := 0; i < 64; i++ {
		square[i] = i * 2
		if square[i] == 0 {
			square[i] = 1
		}
		sum = sum + square[i]
		fmt.Printf("Square: %d\n", i)
		fmt.Printf("Mount for square: %d\n", square[i])
		fmt.Printf("Sum of squares: %d\n\n", sum)
	}
}
