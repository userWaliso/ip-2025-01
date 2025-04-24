
package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Digite o número")
	fmt.Scan(&n)
	
	if n > 0 {
		fmt.Println("POSITIVO")
	} else if n == 0 {
		fmt.Println("NULO")
	} else if n < 0 {
		fmt.Println("NEGATIVO")
	}
}
