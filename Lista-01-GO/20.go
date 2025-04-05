package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, t int

	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)

	t = n1*60*60 + n2*60 + n3
	fmt.Printf("O TEMPO EM SEGUNDOS É = %d\n", t)
}
