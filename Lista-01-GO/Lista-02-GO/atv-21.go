package main
import (
	"fmt"
)

func main() {
	var (
		n int
		nota1, nota2, nota3, exec float64
		conceito string
	)

	fmt.Println("nº do aluno, nota das 3 provas e média dos exercícios")
	fmt.Scan(&n, &nota1, &nota2, &nota3, &exec)

	mediaFinal := ( nota1 + nota2 * 2 + nota3 * 3 + exec) / 7

	if mediaFinal >= 0 && mediaFinal < 4 {
		conceito = "E"
	} else if mediaFinal >= 4 && mediaFinal < 6 {
		conceito = "D"
	} else if mediaFinal >= 6 && mediaFinal < 7.5 {
		conceito = "C"
	} else if mediaFinal >= 7.5 && mediaFinal < 9 {
		conceito = "B"
	} else if mediaFinal >= 9 && mediaFinal <= 10 {
		conceito = "A"
	}
	fmt.Printf("N° do aluno: %d\nNota da prova 1: %.1f\nNota da prova 2: %.1f\nNota da prova 3: %.1f\nMedia dos exercícios: %.1f\nMedia final: %.1f\nConceito: %s\n", n, nota1, nota2, nota3, exec, mediaFinal, conceito)
	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}