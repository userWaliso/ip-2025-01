package main

import (
	f "fmt"
)

func fatorial(a float64) float64 {
	var fat float64 = 1
	for i := 1.0; i <= a; i++ {
		fat *= i
	}
	return fat
}

func main() {
	var x, s float64

	f.Scan(&x)
	s = x
	for i := 1; i <= 20.0; i++ {
		termo := x / fatorial(float64(i))

		if i%2 == 0 {
			s += termo
		} else {
			s -= termo
		}
	}
	f.Println(s)
}
