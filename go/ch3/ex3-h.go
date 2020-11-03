package main

import (
	"fmt"
)

func main() {
	var length, width, height float64

	fmt.Print("Enter Length: ")
	fmt.Scanf("%f", &length)
	fmt.Print("Enter width: ")
	fmt.Scanf("%f", &width)
	fmt.Print("Enter height: ")
	fmt.Scanf("%f", &height)

	volume := (length * width * height)

	fmt.Printf("The volume is: %g", volume)
}
