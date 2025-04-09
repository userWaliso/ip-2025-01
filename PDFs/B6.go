package main

import (
	"fmt"
)

func main() {
	var conta int
	var credito, saldo, depo, reti, valor float64
	fmt.Scan(&conta, &credito, &saldo, &depo, &reti)

	valor = saldo + depo - reti
	if valor <= credito {
		fmt.Printf("Saldo: %.5f\n", valor)
	} else if valor > credito {
		fmt.Printf("Credito excedido: %.5f\n", valor)
	}
}
