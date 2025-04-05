package main

import (
	"fmt"
)

func main() {
	var nota float64
	var conceito string

	fmt.Scan(&nota)

	if nota >= 0.0 && nota <= 6.0 {
		conceito = "D"
	} else if nota >= 6.1 && nota <= 7.5 {
		conceito = "C"
	} else if nota >= 7.6 && nota <= 9.0 {
		conceito = "B"
	} else if nota >= 9.1 && nota <= 10.0 {
		conceito = "A"
	}
	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}
