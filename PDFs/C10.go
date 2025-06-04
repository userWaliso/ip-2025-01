package main

import "fmt"

func main() {
	var P_inicial, D, DeltaP, P_min float64
	var Q_inicial, DeltaQ int
	fmt.Scan(&P_inicial, &Q_inicial, &D, &DeltaP, &DeltaQ, &P_min)

	maxLucro := -1.0
	precoMax := 0.0
	ingressosMax := 0

	Q := Q_inicial
	P := P_inicial

	fmt.Println("Preco    Ingressos Vendidos   Lucro")

	for {
		// Calcular lucro para o preço atual
		lucro := (P * float64(Q)) - D

		// Atualizar máximo
		if lucro > maxLucro {
			maxLucro = lucro
			precoMax = P
			ingressosMax = Q
		}

		// Imprimir linha
		fmt.Printf("%.2f    %d    %.2f\n", P, Q, lucro)

		// Verificar se o próximo preço será válido
		novoP := P - DeltaP
		if novoP < P_min {
			break
		}

		// Atualizar para a próxima iteração
		P = novoP
		Q += DeltaQ
	}

	fmt.Printf("Lucro maximo: %.2f\n", maxLucro)
	fmt.Printf("Na faixa de preco: %.2f com %d ingressos.\n", precoMax, ingressosMax)
}
