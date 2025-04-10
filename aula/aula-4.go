package main

import (
	"fmt"
)
func maior(a, b, c int) int {
	var maior int
	for {
	if a > maior {
		maior = a
	} else if b > maior {	
		maior = b
	} else if c > maior {
		maior = c
	} else {

		break	
	}
		} 
	return maior
}
func main() {

	fmt.Printf("%d\n", maior(1, 2, 3))
	}
