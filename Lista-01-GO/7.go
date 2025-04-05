package main

import (
	"fmt"
	"math"
)

func main() {
	var t, m, c, mm float64

	fmt.Print("Digite a temperatura em F°\n")
	fmt.Scan(&t)
	fmt.Print("Digite a distância em polegadas\n")
	fmt.Scan(&m)

	c = (5*t - 160) / 9
	mm = m * 25.4

	fmt.Printf("VALOR EM CELSIUS = %.2F\nQUANTIDADE DE CHUVA = %.2F\n", math.Round(c), math.Round(mm))
}
