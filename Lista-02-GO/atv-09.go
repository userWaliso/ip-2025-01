package main

import "fmt"

func main() {
	var valor, vendas float64
	
	fmt.Scan(&valor)
	 if valor < 10 {
		vendas = valor * 1.70
	 } else if valor >= 10 && valor <= 30 {
		vendas = valor * 1.50
	 } else if valor > 30 && valor <= 50 {
		vendas = valor * 1.40
	 } else if valor > 50{
		vendas = valor * 1.30
	 }
	 fmt.Printf("%.2f\n", vendas)
}