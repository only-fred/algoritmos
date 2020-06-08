package main

import "fmt"

func main() {
	var value, media int

	i := 0

	for value >= 0 {
		fmt.Print("Value: ")
		fmt.Scanf("%d", &value)

		if value >= 0 {
			media = media + value
			i++
		}
	}

	media = media / i
	fmt.Printf("Media: %d", media)

}
