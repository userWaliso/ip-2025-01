package main

import (
	"fmt"
)

func main() {
	var x, y, z int

	fmt.Scan(&x, &y, &z)

	if x+y <= z || x+z <= y || z+y <= x {
		fmt.Printf("Nao forma triangulo")
	} else if
}
