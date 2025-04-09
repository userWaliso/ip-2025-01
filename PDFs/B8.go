package main

import (
	"fmt"
)

func main() {
	var sal, vendas, com float64

	fmt.Scan(&vendas)
	com = vendas * 0.09

	if vendas > 15000 {
		sal = (500 + com) + 800
		fmt.Printf("%.5f\n", sal)
	} else {
		sal = 500 + com
		fmt.Printf("%.5f\n", sal)
	}
}
