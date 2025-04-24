package main

import (
	"fmt"
)

func main() {
	var conta, valor, consumo float64
	var tipo string

	fmt.Println("Escreva a conta, o consumo e o tipo:")
	fmt.Scan(&conta, &consumo, &tipo)

	if tipo == "r" {
		valor = 5 + consumo*0.05
	} else if tipo == "c" {
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (consumo-80)*0.25
		}
	} else if tipo == "i" {
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (consumo-100)*0.04
		}
	}
	fmt.Printf("CONTA = %.0f\nVALOR A SER PAGO = R$ %.2f\n", conta, valor)
}
