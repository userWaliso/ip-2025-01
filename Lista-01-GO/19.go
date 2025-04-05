package main

import (
	"fmt"
)

func main() {
	var n int
	var soma float64 = 0
	fmt.Scan(&n)

	if n == int(n) && n > 1 {
		for i := 1; i <= n; i++ {
			soma += 1.0 / float64(i)
		}
		fmt.Printf("%.6f\n", soma)
	} else {
		fmt.Print("NÚMERO INVÁLIDO")
	}
}
