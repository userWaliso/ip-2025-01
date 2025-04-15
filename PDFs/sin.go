package main

import (
	"fmt"
	"math"
)

func fatorial(n int) float64 {
	fat := 1.0
	for i := 2; i <= n; i++ {
		fat *= float64(i)
	}
	return fat
}

func main() {
	var x float64
	fmt.Scan(&x)

	for a := 0.0; a <= x+1e-9; a += 0.1 { // soma 1e-9 para evitar erro de precisão
		sin := a
		sin -= math.Pow(a, 3)/fatorial(3)
		sin += math.Pow(a, 5)/fatorial(5)
		sin -= math.Pow(a, 7)/fatorial(7)
		fmt.Printf("%.1f %.4f\n", a, sin)
	}
}
