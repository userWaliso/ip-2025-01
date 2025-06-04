package main

import (
	"fmt"
)

func main() {
	var (
		vet1 [10]int
		vet2 [5]int
	)

	for i := 0; i < 10; i++ {
		fmt.Scan(&vet1[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Scan(&vet2[i])
	}
	SS := 0
	for _, num := range vet2 {
		SS += num
	}

	var pares []int
	var impares []int

	for _, num := range vet1 {
		if num%2 == 0 {
			pares = append(pares, num+SS)
		} else {
			impares = append(impares, num+SS)
		}
	}
	fmt.Println("Vetor resultante par:", pares)
	fmt.Println("Vetor resultante ímpar:", impares)
}
