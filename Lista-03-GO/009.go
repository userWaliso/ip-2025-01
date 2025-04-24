package main
import (
	f"fmt"
)
func main() {
	var (
		x, apv, rpv, exa int
		nota1, nota2 float64
	)
	conceito := []string{}
	mediaGeral := []float64{}
	f.Scan(&x)

	for i:=0;i<x;i++ {
		f.Scan(&nota1, &nota2)
		media := (nota1 + nota2) / 2
		if media >= 0.0 && media <= 3.0 {
			conceito = append(conceito, "Reprovado")
			rpv++
		} else if media > 3.0 && media < 7.0 {
			conceito = append(conceito, "Exame")
			exa++
		} else if media >= 7.0 && media <= 10.0 {
			conceito = append(conceito, "Aprovado")
			apv++
		} 
		mediaGeral = append(mediaGeral, media)
		
	}
	soma := 0.0
	for _, v := range mediaGeral {
		soma += v
	}
	mediaTotal := soma / float64(x)
	for i:=0;i<x;i++ {
	f.Printf("Aluno %d: %.1f. Situação: %s\n", i, mediaGeral[i], conceito[i])	
}
	f.Printf("Total Aprovados: %d\nTotal Exame: %d\nTotal Reprovados: %d\n", apv, exa, rpv)
	f.Printf("Media Geral: %.1f\n", mediaTotal)	
}