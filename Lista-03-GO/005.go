package main
import (
	f"fmt"
)

func main() {
var (
	y, idd, iddcount, altcount, pesocount, pessoas int
	alt, peso, pesoPer, media float64
)
	minAlt := []float64{}
	x := 1
	for i:=0;i<x;i++ {
		f.Scan(&idd, &alt, &peso)
		if idd > 50 {
			iddcount++
		}
		if idd >=10 && idd <= 20 {
			altcount++
			minAlt = append(minAlt, alt)
		}
		if peso < 40 {
			pesocount++
		}
		pessoas++
		f.Print("Deseja continuar?\n")
		f.Scan(&y)
		if y != 1 {
			break
		} else {
			x++
		}
	}
	soma := 0.0
	for _, v := range minAlt {
		soma += v
	}
	media = soma / float64(altcount)
	pesoPer = (float64(pesocount) * 100) / float64(pessoas) 

	f.Printf("Quantidade de pessoas com mais de 50 anos: %d\n", iddcount)
	f.Printf("Média das alturas as pessoas com idade entre 10 e 20 anos: %.2f\n", media)
	f.Printf("Porcentagem de pessoas com menos de 40kg: %.2f%%\n", pesoPer)
}