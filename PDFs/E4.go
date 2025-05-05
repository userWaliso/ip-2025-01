package main

import (
	f "fmt"
	m "math"
)

func main() {
	var n int
	f.Scan(&n)

	vetor := make([]float64, n)
	for i := 0; i < n; i++ {
		f.Scan(&vetor[i])
	}

	s := 0.0
	metade := n / 2

	for i := 0; i < metade; i++ {
		j := n - 1 - i
		diferenca := vetor[i] - vetor[j]
		s += m.Pow(diferenca, 3)
	}

	f.Printf("%.2f\n", s)
}