package main

import "fmt"

func main() {
	var dist, time float64

	fmt.Print("Type distance: ")
	fmt.Scanf("%g", &dist)
	fmt.Print("Type time: ")
	fmt.Scanf("%g", &time)

	velocity := (dist * 1000) / (time * 60)

	fmt.Printf("The velocity is: %g", velocity)
}
