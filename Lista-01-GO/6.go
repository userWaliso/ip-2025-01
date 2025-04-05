package main

import (
	"fmt"
	"math"
)

func main() {

	var C, F float64
	var n int

	fmt.Println("Digite o número de temperaturas a serem calculadas: ")
	fmt.Scan(&n)

	vet := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Print("Digite a temperatura em Fahrenheint")
		fmt.Scan(&F)
		C = (9 / 5) * (F - 32)
		vet[i] = C
	}
	for i := 0; i < n; i++ {
		F = vet[i]*9.0/5.0 + 32
		fmt.Printf("%.2f Fahrenheit equivale a %.2f Celsius\n", F, math.Round(vet[i]*100)/100)
	}
}
