package main

import (
	"fmt"
)
func main() {
	var  (
		n int
		salBruto, salLi, hs, horas, desconto float64
	)
	const (
		sal = 788
		hre = 10
)
	fmt.Println("id do trabalhador e horas extras trabalhadas")
	fmt.Scan(&n, &horas)

	hs = horas * hre
	salBruto = 3 * sal + hs

	if salBruto > 2000 {
		salLi = salBruto * 0.80
		desconto = (salLi - salBruto)*(-1)
	} else {
		salLi = salBruto * 0.88
		desconto = (salLi - salBruto)*(-1)
	}
	fmt.Printf("id:%d\nSalário bruto R$:%.2f\nSalário liquido R$:%.2f\nDesconto: R$:%.2f\n",n ,salBruto, salLi, desconto)
}