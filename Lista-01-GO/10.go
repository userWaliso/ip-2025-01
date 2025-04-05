package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, valor float64

	fmt.Print("Defina A, B, C e D\n")
	fmt.Scan(&a, &b, &c, &d)

	valor = (a * d) - (b * c)

	fmt.Printf("O VALOR DA DETERMINANTE É = %.2f\n", valor)
}
