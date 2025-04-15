package main

import (
	"fmt"
	"math"
)

func fatorial(a, fat int) int {
	for i := 2; i <= a; i++ {
		fat *= i
	}
	return fat
}

func main() {
	var (
		x, cos float64
	)
	fmt.Scan(&x)
	for i := 0; i <= 20; i++ {
		n := 2 * i
		pot := float64(math.Pow(x, float64(n)))
		termo := pot / float64(fatorial(n, 1))
		if i%2 == 0 {
			cos += termo
		} else {
			cos -= termo
		}
	}
	cosFunc := math.Cos(x)
	dif := cos - cosFunc
	fmt.Printf("%.4f %.4f %.4f\n", cos, cosFunc, dif)
}
