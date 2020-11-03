package main

import "fmt"

func main() {
	var value [10]float64
	var sum float64

	sum = 0

	for i := 0; i < 10; i++ {
		fmt.Print("Type you number: ")
		fmt.Scanf("%f", &value[i])

		sum = sum + value[i]
	}

	sum = sum / 10

	fmt.Printf("The avarage is: %f", sum)
}
