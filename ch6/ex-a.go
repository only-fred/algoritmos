package main

import "fmt"

func main() {
	var name [10]string

	for i := 0; i < 10; i++ {
		fmt.Print("Name: ")
		fmt.Scanf("%s", &name[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%dº name is: %s\n", i, name[i])
	}
}
