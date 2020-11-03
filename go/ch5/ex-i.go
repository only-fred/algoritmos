package main

import "fmt"

func main() {
	var fib1, fib2, soma = 1, 0, 0
	for i := 3; i <= 50; i++ {
		soma = fib1 + fib2
		fib1 = fib2
		fib2 = soma
		fmt.Print(fib2, "\n")
	}
}
