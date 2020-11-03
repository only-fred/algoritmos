package main

import "fmt"

func main() {
	var number int

	fmt.Print("Type a number: ")
	fmt.Scanf("%d", &number)

	multiplicationTable(number)
}

func multiplicationTable(number int) {
	for i := 1; i <= 10; i++ {
		result := number * i
		fmt.Printf("%d * %d = %d\n", number, i, result)
	}
}
