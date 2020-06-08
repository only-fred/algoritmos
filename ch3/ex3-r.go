package main

import "fmt"

func main() {
	var a, b, c, null, white, total float64

	a, b, c, null, white = 10, 15, 5, 8, 3

	total = total + (a + b + c + null + white)

	percentA := a * (total / 100)
	percentB := b * (total / 100)
	percentC := c * (total / 100)
	percentNull := null * (total / 100)
	percentWhite := white * (total / 100)

	fmt.Printf("A: %g\nB: %g\nC: %g\nNull: %g\nWhite: %g", percentA, percentB, percentC, percentNull, percentWhite)
}
