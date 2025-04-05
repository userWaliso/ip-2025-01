package main

import (
	"fmt"
	"math"
)

func main() {
	var r, h, lat, area, base, valor float64
	const pi = 3.14159

	fmt.Print("Valor do raio = ")
	fmt.Scan(&r)
	fmt.Print("Valor da altura = ")
	fmt.Scan(&h)

	base = pi * math.Pow(r, 2) * 2
	lat = 2 * pi * r * h
	area = base + lat
	valor = area * 100

	fmt.Printf("O VALOR DO CUSTO É = %.2f\n", valor)
}
