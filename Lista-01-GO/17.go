package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	if x%2 != 0 {
		fmt.Print("O PRIMEIRO NUMERO NAO E PAR")
	} else {
		for i := 0; i < y; i++ {
			fmt.Printf("%d ", x+(i*2))
		}
	}
}
