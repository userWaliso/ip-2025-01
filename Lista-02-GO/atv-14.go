package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	var car float64
	fmt.Println("Valor do carro")
	fmt.Scan(&car)
	
	fmt.Println("Ar Condicionado?")
	fmt.Scan(&a)
	if a == 1 {
		car += 1750
	}
	fmt.Println("Pintura metálica?")
	fmt.Scan(&b)
	if b == 1 {
		car += 800
	}
	fmt.Println("Vidro elétrico?")
	fmt.Scan(&c)
	if c == 1 {
		car += 1200
	}
	fmt.Println("Direção hidráulica?")
	fmt.Scan(&d)
	if d == 1 {
		car += 2000
	}
	fmt.Printf("o valor do carro será de R$:%.2f\n", car)
}