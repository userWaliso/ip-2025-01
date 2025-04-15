package main

import (
	"fmt"
)

func main() {
	var (
		alunos, apv, rpv, exe int
	)
		

	fmt.Scan(&alunos)

	notas := make([]struct {
			nota1 float64
			nota2 float64
			media float64
			conceito string
		}, alunos)

	for i:=0;i<alunos;i++ {
		fmt.Scan(&notas[i].nota1, &notas[i].nota2)
		notas[i].media = (notas[i].nota1 + notas[i].nota2) / 2


		if notas[i].media >=0 && notas[i].media <= 3 {
			notas[i].conceito = "Reprovado"
		} else if notas[i].media > 3 && notas[i].media < 7 {
			notas[i].conceito = "Exame"
		} else if notas[i].media >= 7 && notas[i].media <= 10 {
		 	notas[i].conceito = "Aprovado"
		}
			if notas[i].conceito == "Reprovado" {
				rpv++
			} else if notas[i].conceito == "Aprovado" {
				apv++
			} else if notas[i].conceito == "Exame" {
				exe++
			}
	}

	for  i:=0;i<alunos;i++ {
		fmt.Printf("Aluno %d: %s\n", i+1, notas[i].conceito)
	}
	fmt.Printf("Total Aprovados: %d\nTotal Exame: %d\nTotal Reprovados: %d\n", apv, exe, rpv)
}
