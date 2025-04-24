package main

import (
	f"fmt"
)

func main() {
	var x int

	f.Scan(&x)

	if x <= 0 {
		f.Print("Número inválido")
		return
	}
	fat := 1
	for i := 1; i <= x; i++ {
		fat *= i
		f.Printf("%d ", fat)
	}
}
