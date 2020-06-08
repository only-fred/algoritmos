package main

import "fmt"

func main() {
	var name string
	var gen string

	fmt.Print("Name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Gender(M or F): ")
	fmt.Scanf("%s", &gen)

	if gen == "M" {
		fmt.Printf("Greatings Mr. %s", name)
	} else if gen == "F" {
		fmt.Printf("Greatings Mra. %s", name)
	} else {
		fmt.Print("Invalid gender")
	}
}
