package main

import (
	f "fmt"
)
func main() {
	var vet = make([]int, 10)

	for i:=0;i<10;i++ {
		f.Scan(&vet[i])	
	}
	n := 0
	min := vet[0]
	for i,v := range vet {
		if v < min {
			min = v
			n = i
		}
	}
	f.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", min, n)
}