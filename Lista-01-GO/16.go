package main

import (
	"fmt"
)

func main() {
	var sal, r float64

	fmt.Scan(&sal)

	if sal <= 300 {
		r = sal * 1.5
	} else if sal > 300 {
		r = sal * 1.3
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", r)
}
