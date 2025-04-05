package main

import (
	"fmt"
)

func main() {
	var h, valor, inteiro, sobra int

	fmt.Scan(&h)

	inteiro = h / 3
	sobra = h % 3
	valor = inteiro*10 + sobra*5
	fmt.Printf("O VALOR A PAGAR É = %.2f\n", float64(valor))
}
