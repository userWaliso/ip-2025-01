package main

import (
	"fmt"
)

func main() {
	var n, h, b int

	fmt.Scan(&n)
	
	b = n / 10
	h = b%10
	if n > 99 && n < 1000 {
		fmt.Println(h)
	} else {
		fmt.Println("Numero invalido")
	}
}