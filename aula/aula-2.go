package main

import (
	"fmt"
)

func main() {
	var alunos int
	var soma float64

	fmt.Print("Escreva a quantidade de alunos\n")
	fmt.Scan(&alunos)

	notas := make([]float64, alunos)

	for i := 0; i < alunos; i++ {
		fmt.Printf("Digite a nota do aluno %d°: ", i+1)
		fmt.Scan(&notas[i])
		soma += notas[i]
	}
	fmt.Printf("A soma da notas será: %.2f\n", soma)
}
