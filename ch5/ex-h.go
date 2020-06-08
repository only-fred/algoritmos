package main

import "fmt"

func main() {
	var base, expoent int

	fmt.Print("Type base: ")
	fmt.Scanf("%d", &base)
	fmt.Print("Type expoent: ")
	fmt.Scanf("%d", &expoent)

	expoentFunc(base, expoent)
}

func expoentFunc(base, expoent int) {
	test := 1
	expo := 0

	for i := 0; i <= expoent; i++ {
		result := test
		expo = test
		test = expo * base
		fmt.Print(result, "\n")
	}
}
