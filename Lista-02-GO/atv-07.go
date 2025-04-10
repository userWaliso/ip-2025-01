package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var MAIOR, INTER, MENOR int

	fmt.Scan(&a, &b, &c)

	if a > b && a > c && b < c {
		MAIOR = a
		INTER = c
		MENOR = b
	} else if a > b && a > c && c < b {
		MAIOR = a
		INTER = b
		MENOR = c
	}
	if b > a && b > c && a < c {
		MAIOR = b
		INTER = c
		MENOR = a
	}
	if b > a && b > c && c < a {	
		MAIOR = b
		INTER = a
		MENOR = c
	}
	if c > a && c > b && a < b {
		MAIOR = c
		INTER = b
		MENOR = a
	}
	if c > a && c > b && b < a {
		MAIOR = c
		INTER = a
		MENOR = b
	}
	fmt.Printf("%d %d %d\n", MENOR, INTER, MAIOR)
}
