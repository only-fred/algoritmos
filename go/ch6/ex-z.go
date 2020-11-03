package main

import "fmt"

func main() {
	var a [10]int
	var odd int
	var percent float64

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%d", &a[i])

		if (a[i] % 3) == 0 {
			odd = odd + 1
		}
	}

	percent = float64(len(a)) * (float64(odd))
	fmt.Printf("%d odd, %g%", odd, float64(percent))
}
