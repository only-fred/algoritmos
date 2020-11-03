package main

import "fmt"

func main() {
	var value, sum [15]int

	for i := 0; i < 15; i++ {
		fmt.Printf("Type the %d number? ", i+1)
		fmt.Scanf("%d", &value[i])

		sum[i] = 1

		for j := 0; j < value[i]; j++ {
			sum[i] = sum[i] * (value[i] - j)
		}
	}

	for i := 0; i < 15; i++ {

		fmt.Printf("%d! = %d\n", value[i], sum[i])
	}
}
