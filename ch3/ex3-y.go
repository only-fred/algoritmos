package main

import "fmt"

func main() {
	var a int

	fmt.Print("Type A: ")
	fmt.Scanf("%d", &a)

	antecedent := a - 1
	sucessor := a + 1

	fmt.Printf("Antecedent: %d\nSucessor: %d", antecedent, sucessor)
}
