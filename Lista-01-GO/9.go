package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, delta float64

	fmt.Print("Escreva o valor de A, B e C = ")
	fmt.Scan(&a, &b, &c)

	delta = math.Pow(b, 2) - 4*a*c

	fmt.Printf("O VALOR DE DELTA É = %.2f", delta, "\n")
}
