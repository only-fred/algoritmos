package main

import "fmt"

func main() {
	var a [6]float64
	var b [6]float64

	for i := 0; i < len(a); i++ {
		fmt.Print("Type A: ")
		fmt.Scanf("%g", &a[i])
	}
	for i := 0; i < len(a); i++ {
		if (i % 2) == 0 {
			b[i] = a[(i + 1)]
		} else {
			b[i] = a[(i - 1)]
		}
		fmt.Printf("%g", b[i])
	}
}

// para I de 1 até 6 passo 1 faça
//  R ← I – 2 * (I div 2)
//  se (R <> 0) então
//  B[I] ← A[I + 1]
//  senão
//  B[I] ← A[I - 1]
//  fim_se
//  fim_para
