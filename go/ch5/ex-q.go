package main

import "fmt"

func main() {
	var width, lenght, sum float64
	var name, cont string

	for cont != "NO" {
		fmt.Print("Type name of room: ")
		fmt.Scanf("%s", &name)

		fmt.Print("Type width: ")
		fmt.Scanf("%f", &width)

		fmt.Print("Type length: ")
		fmt.Scanf("%f", &lenght)

		sum = sum + (width * lenght)

		fmt.Println("Do you want stop? \n(ENTER to continue / Type NO to stop)")
		fmt.Scanf("%s", &cont)

		if cont == "no" || cont == "No" {
			cont = "NO"
		}
	}
	fmt.Printf("%fm²", sum)

}
