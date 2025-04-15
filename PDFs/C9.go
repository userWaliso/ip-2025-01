package main

import (
	"fmt"
	"math"
)

func fatorial(a, fat float64) float64 {
	for i := 2.0; i <= a; i++ {
		fat *= i
	}
	return fat
}

func main() {
	var (
		x, cos float64
	)
	fmt.Scan(&x)
	for i := 0.0; i <= 20; i++ {
		n := 2.0 * i
		pot := float64(math.Pow(x, float64(n)))
		termo := pot / fatorial(n, 1)
		if int(i)%2 == 0 {
			cos += termo
		} else {
			cos -= termo
		}
	}
	cosFunc := math.Cos(x)
	dif := cos - cosFunc
	fmt.Printf("%.5f %.5f %.5f\n", cos, cosFunc, dif)
}
