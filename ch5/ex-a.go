package main

import (
	"fmt"
	"math"
)

func main() {
	number := 15

	for i := 15; i <= 200; i++ {
		result := math.Pow(float64(number), 2)
		number = number + 1
		fmt.Printf("%d\n", int(result))
	}
}
