package main

import (
	"fmt"
)

func main() {
	var n int
	var popular, people, geral, arqui, chair, valor float64

	fmt.Print("Número de jogos: \n")
	fmt.Scan(&n)

	rendas := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Print("Número de pessoas; Porcentagem de: Popular, Geral, Arquibancada e Cadeiras\n")
		fmt.Scan(&people, &popular, &geral, &arqui, &chair)

		valor = (people * (popular / 100)) +
			((people * (geral / 100)) * 5) +
			((people * (arqui / 100)) * 10) +
			((people * (chair / 100)) * 20)

		rendas[i] = valor
	}
	for i, valor := range rendas {
		fmt.Printf("A RENDA DO JOGO N. %d = R$ %.2f\n", i+1, valor)
	}
}
