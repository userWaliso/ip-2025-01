package main

import (
	"fmt"
)

func main() {
	var n1 int

	fmt.Scan(&n1)

	if n1%5 == 0 {
		fmt.Printf("O NÚMERO É DIVISÍVEL\n")
	} else {
		fmt.Printf("O NÚMERO NÃO É DIVISÍVEL\n")
	}
}
