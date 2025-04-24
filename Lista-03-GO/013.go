package main

import (
	f "fmt"
)

func main() {
	var h, n, soma float64
	for i := 1.0; i <= 50; i++ {
		n = (i * 2) - 1
		dem := i
		h = n / dem
		soma += h
	}
	f.Println(soma)
}